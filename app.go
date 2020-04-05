package main

import (
	"fmt"
	"log"

	"ssh/connection"
)

func main() {
	allConn, err := connection.OpenConnections()
	if err != nil {
		log.Println(err)
	} else {
		for i := range allConn {
			go allConn[i].SendCommands()
		}
	}
	fmt.Println("Enter for finish")
	fmt.Scanln()
}
