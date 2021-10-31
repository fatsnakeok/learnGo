package main

import "net"

/**
1.什么是粘包？
 粘包问题是指当发送两条消息时，比如发送了 123 和 456，但另一端接收到的却是 1234，像这种一次性读取了两条数据的情况就叫做粘包（正常情况应该是一条一条读取的）。

2.什么是半包？
 半包问题是指，当发送的消息是 123 时，另一端却接收到的是 12 和 3 两条信息，像这种情况就叫做半包
3.为什么会有粘包和半包？
  这是因为 TCP 是面向连接的传输协议，TCP 传输的数据是以流的形式，而流数据是没有明确的开始结尾边界，所以 TCP 也没办法判断哪一段流属于一个消息。

4.粘包的主要原因：
  发送方每次写入数据 < 套接字（Socket）缓冲区大小；
  接收方读取套接字（Socket）缓冲区数据不够及时。
5.半包的主要原因：
  发送方每次写入数据 > 套接字（Socket）缓冲区大小；
  发送的数据大于协议的 MTU (Maximum Transmission Unit，最大传输单元)，因此必须拆包。
*/

import (
	"fmt"
)

/**
  fix_length
    即每次发送固定缓冲区大小.客户端和服务器约定每次发送请求的大小.例如客户端发送1024个字节，服务器接受1024个字节。
    这样虽然可以解决粘包的问题，但是如果发送的数据小于1024个字节，就会导致数据内存冗余和浪费；且如果发送请求大于1024字节，会出现半包的问题。
*/
func client_tcp_fix_length(conn net.Conn) {
	fmt.Println("client, fix length")
	sendByte := make([]byte, 1024)
	sendMsg := "{\"test1\":1,\"test2\": 2}"
	for i := 0; i < 1000; i++ {
		tempByte := []byte(sendMsg)
		for j := 0; j < len(tempByte) && j < 1024; j++ {
			sendByte[j] = tempByte[j]
		}
		_, err := conn.Write(sendByte)
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}

/**
	delimiter based
	 基于定界符来判断是不是一个请求（例如结尾'\n'). 客户端发送过来的数据，每次以\n结束，服务器每接受到一个 \n 就以此作为一个请求.
    这种方式的缺点在于如果数据量过大，查找定界符会消耗一些性能
*/
func client_tcp_delimiter(conn net.Conn) {
	fmt.Println("client, delimiter based")
	var sendMsgs string
	sendMsg := "{\"test1\":1,\"test2\":2}\n"
	for i := 0; i < 1000; i++ {
		sendMsgs += sendMsg
		_, err := conn.Write([]byte(sendMsgs))
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}

/**
length field based frame decoder
在TCP协议头里面写入每次发送请求的长度。 客户端在协议头里面带入数据长度，服务器在接收到请求后，根据协议头里面的数据长度来决定接受多少数据。
*/
func client_tcp_frame_decoder(conn net.Conn) {
	fmt.Println("client, length field based frame decoder")
	for i := 0; i < 1000; i++ {
		sendMsg := "{\"test01\":1,\"test02\":2}"
		_, err := conn.Write(([]byte(sendMsg)))
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client_tcp_fix_length(conn)
	// client_tcp_delimiter(conn)
}
