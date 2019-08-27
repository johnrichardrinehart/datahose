package sinks

import (
	"os"
)

type file struct {
	*os.File
}

func newFile(name string) (Sink, error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	return file{f}, err
}
