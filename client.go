package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
	"strings"
)

const (
	END_BYTES = "\000\001\002\003\004\005"
	ADDR_SERVER = ":8080"
)

func main() {

	conn, err := net.Dial("tcp", ADDR_SERVER)
	if err != nil {
		panic("can't connect to server")
	}
	defer conn.Close()
	go ClientOuput(conn)
	ClientInput(conn)

}

func ClientInput(conn net.Conn) {
	for {
		conn.Write([]byte(InputSrting() + END_BYTES))
	}
}

func ClientOuput(conn net.Conn) {
	var (
		buffer = make([]byte, 512)
		message string
	)
	close: for {
		message = ""
		for {
			lenght, err := conn.Read(buffer)
			if lenght == 0 || err != nil { break close }
			message = string(buffer[:lenght])
			if strings.HasSuffix(message, END_BYTES) {
				message = strings.TrimSuffix(message, END_BYTES)
				break
			}
		}
		fmt.Println(message)	
	}
}

func InputSrting() string {
	msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(msg, "\n", "", -1)
}
