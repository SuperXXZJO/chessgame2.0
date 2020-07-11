package game

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var h =Hub{
	Client: make(map[*Client]bool),
	Broadcast: make(chan []byte),
	Register: make(chan *Client),
	Unregister: make(chan *Client),
}
type Hub struct {
	Name string
	Client map[*Client]bool
	Broadcast chan []byte
	Register  chan *Client
	Unregister chan *Client
}



type Client struct {
	Name string
	Conn *websocket.Conn
	Send chan []byte

}

func Startgame(c *gin.Context) {
	name:=c.PostForm("roomname")
	h.Name = name
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//新建一个客户端
	client := &Client{
		Name: c.Query("username"),
		Conn: conn,
		Send: make(chan []byte),
	}
	//
	h.Register<-client
	Newcb()
	go h.start()
	go client.read()
	go client.write()

}

//客户端写数据
func (c *Client)write(){
	for message := range c.Send{
		c.Conn.WriteMessage(websocket.TextMessage,message)

	}
	c.Conn.Close()
}
//客户端读数据
func (c *Client)read(){
	for  {
		_,message,err:=c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var messagejs string
		json.Unmarshal(message,&messagejs)
		Update(&messagejs)
		res,_:=json.Marshal(messagejs)
		h.Broadcast<-res

	}
}

func (h *Hub)start(){
	for  {
		select {
		case client :=<-h.Register:
			h.Client[client] = true
		case client:=<-h.Unregister:
			if _,ok :=h.Client[client];ok{
				delete(h.Client,client)
				close(client.Send)
			}
		case message := <- h.Broadcast :
			for client :=range h.Client {
				select {
				case client.Send <-message:
				default:
					close(client.Send)
					delete(h.Client,client)
				}
			}
		}
	}
}