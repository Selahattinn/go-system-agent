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

	//for parsing flags
	flag.Parse()

	// create walker
	walker := walker.NewWalker(*rootDirectory)

	// create client
	c := client.NewClient("http://"+*serverAdrres+":"+*serverPort+"/api/v1/files", &walker)

	// get informations about files
	err := c.Walk()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// send information to server
	err = c.SendToServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
