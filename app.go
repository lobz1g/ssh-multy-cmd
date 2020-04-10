package main

import (
	"fmt"
	"log"

	"ssh/connection"
	"ssh/message"
)

func sendCmd(connections []*connection.Connection, ch chan bool) {
	for _, c := range connections {
		go func(c *connection.Connection, ch chan bool) {
			msg := c.SendCommands()
			err := msg.Save()
			if err != nil {
				message.NewMessage("localhost", err.Error(), true).Save()
				ch <- false
			} else {
				if msg.IsError {
					ch <- false
				} else {
					ch <- true
				}
			}
			return
		}(c, ch)
	}
}

func main() {
	ch := make(chan bool)
	allConn, err := connection.OpenConnections()
	if err != nil {
		log.Println(err)
	} else {
		sendCmd(allConn, ch)
		for range allConn {
			answ := <-ch
			if answ {
				print("V")
			} else {
				print("X")
			}
		}
	}
	fmt.Println("\nFinished. For close tap enter")
	fmt.Scanln()
}
