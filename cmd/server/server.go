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
	flag.Parse()
	config := &server.Config{
		ListenAddress: *listenAddres,
	}
	i := server.NewInstance(config)
	logrus.Infof("Starting go-system-agent-server %s", version.Info())
	logrus.Infof("Build context %s", version.BuildContext())
	i.Start()
}
