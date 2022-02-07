package api

import (
	"net/http"

	"github.com/Selahattinn/go-system-agent/pkg/server/api/response"
	"github.com/gorilla/mux"
)

// API represents the structure of the API
type API struct {
	Router *mux.Router
}

// New returns the api settings
func New(router *mux.Router) (*API, error) {
	api := &API{
		Router: router,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	//Get All Files Information
	api.Router.HandleFunc("/api/v1/files", api.corsMiddleware(api.logMiddleware(api.getAllFilesInformations))).Methods("POST")
	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

	return
}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
