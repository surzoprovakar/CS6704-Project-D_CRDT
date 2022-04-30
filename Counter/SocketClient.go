package main

import (
	"fmt"
	"net"
	"os"
)

func EstablishConnections(addresses []string) []net.Conn {
	conns := make([]net.Conn, len(addresses))

	for i := 0; i < len(addresses); i++ {
		var err error
		fmt.Println("establishing connection " + addresses[i])
		conns[i], err = net.Dial("tcp", addresses[i])
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			os.Exit(1)
		}
	}
	return conns
}

func Broadcast(conns []net.Conn, content []byte) {

	for i := 0; i < len(conns); i++ {
		_, err := conns[i].Write(content)
		if err != nil {
			fmt.Println("Error socket writing:", err.Error())
			os.Exit(1)
		}
	}
}
