package main

import (
	"fmt"
	"net"
)

func main() {
	mal := make(chan bool)
	flag := true
	client, err := net.Dial("tcp", "127.0.0.8:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("연결된거같음")
	defer client.Close()
	go func(con net.Conn) { //쓰기
		var s string
		data := make([]byte, 4096)
		for {
			fmt.Scanln(&s)
			if s == "done" {
				mal <- true
			}
			data = []byte("")
			data = append([]byte(s), data...)
			_, err := con.Write(data)
			flag = true
			s = ""
			if err != nil {
				fmt.Println(err)
				return
			}
			s = ""
		}

	}(client)

	go func(con net.Conn) { //읽기
		data := make([]byte, 4096)
		for {
			if flag == true {
				n, err := con.Read(data)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(data[:n]))
				flag = false
			}
		}

	}(client)
	<-mal
	return
}
