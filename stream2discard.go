package stream2discard

import (
	"net"
)

type Connection struct{ net.Conn }

type StreamListener struct{ *net.UnixListener }

func (l StreamListener) Close() error { return l.UnixListener.Close() }

func (l StreamListener) Accept() (Connection, error) {
	conn, err := l.UnixListener.Accept()
	return Connection{Conn: conn}, err
}

type StreamConnection struct{ *net.UnixConn }

func (l StreamListener) AcceptUnix() (StreamConnection, error) {
	ucon, err := l.UnixListener.AcceptUnix()
	return StreamConnection{UnixConn: ucon}, err
}

func (c StreamConnection) Close() error { return c.UnixConn.Close() }

func (c StreamConnection) Read(buf []byte) (int, error) {
	return c.UnixConn.Read(buf)
}

type StreamAddr struct{ *net.UnixAddr }

func (a StreamAddr) Listen() (StreamListener, error) {
	listener, err := net.ListenUnix("unix", a.UnixAddr)
	return StreamListener{UnixListener: listener}, err
}

type UnixSockPath string

func (p UnixSockPath) ToAddr() StreamAddr {
	return StreamAddr{
		UnixAddr: &net.UnixAddr{
			Name: string(p),
			Net:  "unix",
		},
	}
}
