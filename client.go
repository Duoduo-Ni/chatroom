package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//main调用完毕，关闭连接
	defer conn.Close()

	//新建协程，从键盘输入内容，给服务器发送内容，任务2
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) //从键盘读取内容
			if err != nil {
				fmt.Println("os.Stdin.errr=", err)
				return
			}
			//把输入的内容给服务器发送
			conn.Write(str[:n])
		}
	}()

	//接收服务器回复的数据，任务1
	//切片缓冲
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收服务器的请求
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println(string(buf[:n])) //打印接收到的内容
	}

}
