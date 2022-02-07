package client

import (
	"testing"

	"github.com/Selahattinn/go-system-agent/pkg/model"
	"github.com/Selahattinn/go-system-agent/pkg/walker"
)

func TestNewClient(t *testing.T) {
	testWalker := walker.NewWalker("./")
	testWalker2 := walker.NewWalker("./")

	type args struct {
		ServerAddress string
		ClientID      int64
		walker        *walker.Walker
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{name: "Create New Client", args: args{ServerAddress: "127.0.0.1", ClientID: 1, walker: &testWalker}, want: &Client{serverAddress: "127.0.0.1", clientID: 1, walker: &testWalker2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.ServerAddress, tt.args.ClientID, tt.args.walker)
			if got.clientID != tt.want.clientID || got.serverAddress != tt.want.serverAddress || got.walker.GetRootDirectory() != tt.want.walker.GetRootDirectory() {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)

			}
		})
	}
}

func TestClient_AddFile(t *testing.T) {
	type fields struct {
		serverAddress string
		clientID      int64
		files         []*model.File
		walker        *walker.Walker
	}
	type args struct {
		file *model.File
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Add File To Files", fields: fields{serverAddress: "127.0.0.1", clientID: 1, files: make([]*model.File, 0), walker: &walker.Walker{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				serverAddress: tt.fields.serverAddress,
				clientID:      tt.fields.clientID,
				files:         tt.fields.files,
				walker:        tt.fields.walker,
			}
			c.AddFile(tt.args.file)
		})
	}
}

func TestClient_ClearFiles(t *testing.T) {
	type fields struct {
		serverAddress string
		clientID      int64
		files         []*model.File
		walker        *walker.Walker
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Clear Files", fields: fields{serverAddress: "127.0.0.1", clientID: 1, files: make([]*model.File, 0), walker: &walker.Walker{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				serverAddress: tt.fields.serverAddress,
				clientID:      tt.fields.clientID,
				files:         tt.fields.files,
				walker:        tt.fields.walker,
			}
			c.ClearFiles()
		})
	}
}
