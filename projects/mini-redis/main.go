package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
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
			reader := bufio.NewReader(c)
			for {
				str, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				cleaned := strings.TrimSpace(str)
				parts := strings.Fields(cleaned)

				if len(parts) == 0 {
					c.Write([]byte("Empty Command." + "\n"))
					continue
				}

				switch parts[0] {
				case "SET":
					if len(parts) < 3 {
						c.Write([]byte("Invalid Input" + "\n"))
						continue
					}
					c.Write([]byte("got SET key=" + parts[1] + " value=" + parts[2] + "\n"))
				case "GET":
					if len(parts) < 2 {
						c.Write([]byte("Invalid Input" + "\n"))
						continue
					}
					c.Write([]byte("got GET key=" + parts[1] + "\n"))
				case "DEL":
					if len(parts) < 2 {
						c.Write([]byte("Invalid Input" + "\n"))
						continue
					}
					c.Write([]byte("got DEL key=" + parts[1] + "\n"))
				default:
					c.Write([]byte("Unrecognized Command." + "\n"))
				}
			}
		}(conn)
	}
}