package main

import (
	"flag"

	"github.com/Selahattinn/go-system-agent/pkg/server"
	"github.com/Selahattinn/go-system-agent/pkg/version"
	"github.com/sirupsen/logrus"
)

var (
	listenAddres = flag.String("listen", "127.0.0.1:8080", "Listen address")
)

func main() {
	// parse flags
	flag.Parse()

	// create server config
	config := &server.Config{
		ListenAddress: *listenAddres,
	}

	// create new server instance
	i := server.NewInstance(config)
	logrus.Infof("Starting go-system-agent-server %s", version.Info())
	logrus.Infof("Build context %s", version.BuildContext())

	//start server
	i.Start()
}
