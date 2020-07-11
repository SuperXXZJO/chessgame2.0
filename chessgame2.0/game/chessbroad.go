package game

import (
	"encoding/json"
	"strconv"
	"strings"
)

const MaxRow = 5  //行
const MaxLine = 5 //列
var chessbroad [MaxRow][MaxLine]string

var x,y int
var name string
//生成一个棋盘
func Newcb (){
	//初始化一个棋盘
	for row := 0; row < 5; row++ {
		for line := 0; line < 5; line++ {
			chessbroad[row][line] = "0"
		}
	}

}

//落子
func Update (msg *string){
	var result []byte
	res :=strings.Split(*msg,".")
	point :=res[0]
	x,_:=strconv.Atoi(res[1])
	y,_:=strconv.Atoi(res[1])
	chessbroad[x][y]=point
	res2:=CheckWinner()
	if res2 != " " {
		result,_ =json.Marshal(res2)
	}
	result,_ =json.Marshal(chessbroad)
	*msg = string(result)
}

//判断是否已经下过了

//判断胜负
func CheckWinner() string {
	//判断是否连成一行
	for row := 0; row < 5; row++ {
		if chessbroad[row][0] == chessbroad[row][1] && chessbroad[row][0] == chessbroad[row][2] && chessbroad[row][0] == chessbroad[row][3] && chessbroad[row][0] == chessbroad[row][4] {
			return chessbroad[row][0]

		}
	}
	//判断是否连成一列
	for line := 0; line < 5; line++ {
		if chessbroad[0][line] == chessbroad[1][line] && chessbroad[0][line] == chessbroad[2][line] && chessbroad[0][line] == chessbroad[3][line] && chessbroad[0][line] == chessbroad[4][line] {
			return chessbroad[0][line]

		}
	}
	//判断对角线是否连成一线
	if chessbroad[0][0] == chessbroad[1][1] && chessbroad[0][0] == chessbroad[2][2] && chessbroad[0][0] == chessbroad[3][3] && chessbroad[0][0] == chessbroad[4][4] {
		return chessbroad[0][0]
	}
	if chessbroad[0][4] == chessbroad[1][3] && chessbroad[0][4] == chessbroad[2][2] && chessbroad[0][4] == chessbroad[3][1] && chessbroad[0][4] == chessbroad[4][0] {
		return chessbroad[0][4]
	}
	return " "
}