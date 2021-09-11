package main


import (
	"fmt"
)

func main() {
	sqlCon, err := initConn()
	defer sqlCon.Close() // 延迟执行：在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行
	if err != nil {
		fmt.Println("initSql error:", err)
		fmt.Printf("%+v", err, "\n")
	}

	u, err := queryByParam(sqlCon, "select id,username from user where id = ?", 2)
	if err != nil {
		fmt.Println("query sql error :", err)
		fmt.Printf("%+v", err, "\n")
	}
	fmt.Println(u)
}
