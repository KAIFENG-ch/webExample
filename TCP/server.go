package main

import (
	"bufio"
	"fmt"
	"net"
)

func Process(conn net.Conn) {
	defer conn.Close()
	for true {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("err")
			break
		}
		recv := string(buf[:n])
		fmt.Println("收到的数据：", recv)
		_, err = conn.Write([]byte(recv))
		if err != nil {
			return
		}
	}
}

func main() {
	listen,err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for true {
		conn,err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println(err)
			return
		}
		go Process(conn)
	}
}