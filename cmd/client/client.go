package main

import (
	"fmt"

	"github.com/Selahattinn/go-system-agent/pkg/client"
	"github.com/Selahattinn/go-system-agent/pkg/walker"
)

func main() {
	walker := walker.NewWalker("/home/selo/personal")
	client := client.NewClient("127.0.0.1", 1, &walker)
	err := client.Walk()
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range client.GetAllFiles() {
		fmt.Println(file)
	}
}
