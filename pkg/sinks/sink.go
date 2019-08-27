package sinks

import (
	"fmt"
	"strings"
)

// Sink interface defines the behavior of a sink of data
type Sink interface {
	Write([]byte) (int, error)
	Close() error
}

//StringToSink converts a string to a value that satisfies the Sink interface
func StringToSink(sinkString string) (Sink, error) {
	switch {
	case sinkString == "stdout":
		return &stdout{}, nil
	case strings.HasSuffix(sinkString, ".out"):
		f, err := newFile(sinkString)
		if err != nil {
			return nil, err
		}
		return f, err
	case strings.HasSuffix(sinkString, ".unix"):
		return newUDS(sinkString)
	case strings.HasPrefix(sinkString, "tcp://"):
		return newTCP(sinkString)
	case sinkString == "":
		return nil, fmt.Errorf("empty sink")
	default:
		return nil, fmt.Errorf("invalid sink: %s", sinkString)
	}
}
