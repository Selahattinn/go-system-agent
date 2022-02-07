package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Selahattinn/go-system-agent/pkg/client"
	"github.com/Selahattinn/go-system-agent/pkg/walker"
)

var (
	rootDirectory = flag.String("root", "./", "Path of root directory for walk")
	serverAdrres  = flag.String("server", "127.0.0.1", "Server address")
	serverPort    = flag.String("port", "8080", "Server Port")
)

func main() {
	flag.Parse()
	walker := walker.NewWalker(*rootDirectory)
	client := client.NewClient("http://"+*serverAdrres+":"+*serverPort+"/api/v1/files", &walker)
	err := client.Walk()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = client.SendToServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
