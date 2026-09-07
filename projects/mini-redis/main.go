package main

import (
	"bufio"
	"fmt"
	"net"
)

func main()  {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		go func(c net.Conn) {
			defer c.Close()
			str, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
	
			c.Write([]byte("ECHO: " + str))
		}(conn)
	}
}