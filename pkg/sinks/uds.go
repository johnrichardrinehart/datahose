package sinks

import (
	"fmt"
	"io"
	"net"
	"time"
)

type unixDomainSocket struct {
	conns    []net.Conn
	listener net.Listener
	quit     chan bool
}

func (u *unixDomainSocket) Write(b []byte) (int, error) {
	var n int
	var err error

	// clean up disconnected clients
	for idx, conn := range u.conns {
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		// need to read at least one byte to trigger EOF
		b := make([]byte, 1)
		if _, err := conn.Read(b); err == io.EOF {
			conn.Close()
			u.conns = append(u.conns[:idx], u.conns[idx+1:]...)
		}
	}
	for _, conn := range u.conns {
		// write to those still connected
		n, err = conn.Write(b)
		if err != nil {
			return n, err
		}
	}
	return n, err
}

func (u *unixDomainSocket) Close() error {
	fmt.Println("Closing unix domain socket")
	close(u.quit)
	u.listener.Close()
	return nil
}

func newUDS(name string) (Sink, error) {
	l, err := net.Listen("unix", name)
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
