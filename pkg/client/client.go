package client

import (
	"github.com/Selahattinn/go-system-agent/pkg/model"
	"github.com/Selahattinn/go-system-agent/pkg/walker"
)

type Client struct {
	serverAddress string
	clientID      int64
	files         []*model.File
	walker        *walker.Walker
}

func NewClient(ServerAddress string, ClientID int64, walker *walker.Walker) *Client {
	return &Client{
		serverAddress: ServerAddress,
		clientID:      ClientID,
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

func (c *Client) SendToServer() {

}

func (c *Client) ClearFiles() {
	c.files = make([]*model.File, 0)
}

func (c *Client) GetAllFiles() []*model.File {
	return c.files
}
