package sinks

import "os"

type stdout struct{}

func (s *stdout) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func (s *stdout) Close() error {
	return os.Stdout.Close()
}
