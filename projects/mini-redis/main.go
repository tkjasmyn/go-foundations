package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func main()  {
	var store = make(map[string]string)
	var mu sync.Mutex


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

					key := parts[1]
					value := parts[2]

					mu.Lock()
						store[key] = value
					mu.Unlock()

					c.Write([]byte("+OK" + "\n"))
				case "GET":
					if len(parts) < 2 {
						c.Write([]byte("Invalid Input" + "\n"))
						continue
					}

					key := parts[1]
					
					mu.Lock()
					if value, ok := store[key]; ok {
						c.Write([]byte(value + "\n"))
						} else {
							c.Write([]byte("(nil)" + "\n"))
						}
					mu.Unlock()
				case "DEL":
					if len(parts) < 2 {
						c.Write([]byte("Invalid Input" + "\n"))
						continue
					}
					
					key := parts[1]

					mu.Lock()
					_, ok := store[key]
					if ok {
						delete(store, key)
						c.Write([]byte(":1" + "\n"))
					} else {
						c.Write([]byte(":0" + "\n"))
					}
					mu.Unlock()
				case "KEYS":
					if len(store) == 0 {
						c.Write([]byte("(empty)" + "\n"))
					}
					
					mu.Lock()
					for key := range store {
						c.Write([]byte(key + "\n"))
					}
					mu.Unlock()
				default:
					c.Write([]byte("Unrecognized Command." + "\n"))
				}
			}
		}(conn)
	}
}