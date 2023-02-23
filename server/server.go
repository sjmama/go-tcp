package main

import (
	"fmt"
	"net"
)

func main() {
	flag := true
	link, err := net.Listen("tcp", "127.0.0.8:8080") //link 랑 error
	if err != nil {
		fmt.Println(err)
		return
	}
	defer link.Close()

	fmt.Println("연결된거같음")

	for {
		conn, err := link.Accept() //이제 이 con으로 주고 받고 하면됨
		if flag == false {
			fmt.Println("00000000000")
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer link.Close()

		go requesthandler(conn, &flag)
		fmt.Println("8888888888888")
	}
}

func requesthandler(conn net.Conn, flag *bool) {
	data := make([]byte, 4096)
	for {
		n, err := conn.Read(data) //읽은 바이트 수랑 에러
		if err != nil {

			*flag = false
			fmt.Println(err)
			fmt.Println(*flag)
			return
		}

		fmt.Println(string(data[:n]), n)
		data = append([]byte("server가 보냄:"), data[:n]...)

		_, err = conn.Write(data)
		if err != nil {
			*flag = false
			fmt.Println(err)
			fmt.Println(*flag)
			return
		}
	}

}
