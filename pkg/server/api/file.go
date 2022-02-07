package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/go-system-agent/pkg/model"
	"github.com/Selahattinn/go-system-agent/pkg/server/api/response"
)

func (a *API) getAllFilesInformations(w http.ResponseWriter, r *http.Request) {
	//handle with req
	var fwReq model.Message
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting message info: %v", err), http.StatusBadRequest, "")
		fmt.Println(err)

		return
	}
	if fwReq.Client == "" {
		response.Errorf(w, r, fmt.Errorf("error getting message info: %v", err), http.StatusBadRequest, "")
		fmt.Println(err)
		return
	}
	fmt.Println("User: ", fwReq.Client)
	fmt.Println("Files:")
	for _, file := range fwReq.Files {
		fmt.Println(file)
	}

	response.Write(w, r, "Supriseee :)")
	return
}
