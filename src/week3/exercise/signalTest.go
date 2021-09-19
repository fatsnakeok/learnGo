package main

import (
	"fmt"
	"os"
	"os/signal"
)

//https://zhuanlan.zhihu.com/p/128953024
func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("start..")
	s := <-c
	fmt.Println("End...", s)
}
