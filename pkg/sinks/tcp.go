package sinks

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

type tcp struct {
	conns    []net.Conn
	listener net.Listener
	quit     chan bool
}

func (t *tcp) Write(b []byte) (int, error) {
	var n int
	var err error

	// clean up disconnected clients
	for idx, conn := range t.conns {
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		// need to read at least one byte to trigger EOF
		b := make([]byte, 1)
		if _, err := conn.Read(b); err == io.EOF {
			conn.Close()
			t.conns = append(t.conns[:idx], t.conns[idx+1:]...)
		}
	}
	for _, conn := range t.conns {
		// write to those still connected
		n, err = conn.Write(b)
		if err != nil {
			return n, err
		}
	}
	return n, err
}

func (t *tcp) Close() error {
	fmt.Println("Closing unix domain socket")
	close(t.quit)
	t.listener.Close()
	return nil
}

func newTCP(name string) (Sink, error) {
	hostPort := strings.Split(name, "tcp://")[1]
	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		return nil, err
	}
	u := &unixDomainSocket{listener: l, quit: make(chan bool)}
	go func() {
		for {
			conn, err := u.listener.Accept()
			if err != nil {
				select {
				case <-u.quit:
					return
				default:
				}
				continue
			}
			u.conns = append(u.conns, conn)
		}
	}()
	return u, nil
}
