package main

import (
	"fmt"
	"net"
)

func main() {
	link, err := net.Listen("tcp", "127.0.0.8:8080") //link 랑 error
	if err != nil {
		fmt.Println(err)
		return
	}
	defer link.Close()

	fmt.Println("연결된거같음")

	for {
		conn, err := link.Accept() //이제 이 con으로 주고 받고 하면됨
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer link.Close()

		go requesthandler(link, conn)
		fmt.Println("8888888888888")
	}
}

func requesthandler(link net.Listener, conn net.Conn) {
	data := make([]byte, 4096)
	for {
		n, err := conn.Read(data) //읽은 바이트 수랑 에러
		if err != nil {
			link.Close()
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n]), n)
		data = append([]byte("server가 보냄:"), data[:n]...)

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
