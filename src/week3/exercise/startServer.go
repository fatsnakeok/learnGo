package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world!")
}

func main() {
	//http.HandleFunc("/hello", helloServer)
	//if err := http.ListenAndServe(":8080", nil); err!=nil {
	//	log.Fatal("Server start error: ", err)
	//}

	srv := &http.Server{Addr: ":9090"}
	http.HandleFunc("/hello", helloServer)
	fmt.Println("http server start")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server start error: ", err)
	}

	// 关闭 server
	//srv.Shutdown(context.TODO())
}
