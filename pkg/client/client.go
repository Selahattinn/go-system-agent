package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"

	"github.com/Selahattinn/go-system-agent/pkg/model"
	"github.com/Selahattinn/go-system-agent/pkg/walker"
)

type Client struct {
	serverAddress string
	files         []*model.File
	walker        *walker.Walker
}

func NewClient(ServerAddress string, walker *walker.Walker) *Client {
	return &Client{
		serverAddress: ServerAddress,
		files:         make([]*model.File, 0),
		walker:        walker,
	}
}

func (c *Client) Walk() error {
	files, err := c.walker.Walk()
	if err != nil {
		return err
	}
	c.files = files
	return nil
}

func (c *Client) AddFile(file *model.File) {
	c.files = append(c.files, file)
}

func (c *Client) SendToServer() error {
	username, err := c.getUser()
	if err != nil {
		return err
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"client": username,
		"files":  c.files,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(c.serverAddress, "application/json", responseBody)
	//Handle Error
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(body)
	fmt.Println(sb)
	return nil
}

func (c *Client) ClearFiles() {
	c.files = make([]*model.File, 0)
}

func (c *Client) GetAllFiles() []*model.File {
	return c.files
}

func (c *Client) getUser() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
