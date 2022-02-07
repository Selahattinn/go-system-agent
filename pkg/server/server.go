package server

import (
	"net/http"
	"os"

	"github.com/Selahattinn/go-system-agent/pkg/server/api"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddress string
}

// Instance represents an instance of the server
type Instance struct {
	Config     *Config
	API        *api.API
	httpServer *http.Server
}

// NewInstance returns an new instance of our server
func NewInstance(cfg *Config) *Instance {
	return &Instance{
		Config: cfg,
	}
}

// Start starts the server
func (i *Instance) Start() {
	var err error
	var router = mux.NewRouter()

	// Initialize API
	i.API, err = api.New(router)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create API instance")
	}
	// Startup the HTTP Server in a way that we can gracefully shut it down again
	i.httpServer = &http.Server{
		Addr:    i.Config.ListenAddress,
		Handler: router,
	}

	err = i.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("HTTP Server stopped unexpected")
		os.Exit(1)
	} else {
		logrus.WithError(err).Info("HTTP Server stopped")
		os.Exit(1)
	}
}
