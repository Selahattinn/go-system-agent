package model

type File struct {
	Path       string
	Content    string
	Size       int64
	Permision  string
	Owner      string
	Name       string
	Isdir      bool
	LastUptade int64
}
