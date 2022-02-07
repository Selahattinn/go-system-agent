package walker

import (
	"testing"

	"github.com/Selahattinn/go-system-agent/pkg/model"
)

func TestNewWalker(t *testing.T) {
	type args struct {
		rootDirectory string
	}
	tests := []struct {
		name string
		args args
		want Walker
	}{
		{name: "Create New Walker", args: args{rootDirectory: "./"}, want: Walker{rootDirectory: "./"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewWalker(tt.args.rootDirectory)
			if got.rootDirectory != tt.want.rootDirectory {
				t.Errorf("NewWalker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalker_Walk(t *testing.T) {
	type fields struct {
		rootDirectory string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.File
		wantErr bool
	}{
		{name: "Walk", wantErr: false, fields: fields{rootDirectory: "./"}, want: []*model.File{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Walker{
				rootDirectory: tt.fields.rootDirectory,
			}
			_, err := w.Walk()
			if (err != nil) != tt.wantErr {
				t.Errorf("Walker.Walk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestWalker_GetRootDirectory(t *testing.T) {
	type fields struct {
		rootDirectory string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Get Root Directory", fields: fields{rootDirectory: "./"}, want: "./"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Walker{
				rootDirectory: tt.fields.rootDirectory,
			}
			if got := w.GetRootDirectory(); got != tt.want {
				t.Errorf("Walker.GetRootDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
