package main

import (
	"context"
	"errors"
	"io"
	"log"
	"os"

	ud "github.com/takanoriyanagitani/go-ustream2discard"
)

var path2unixSock ud.UnixSockPath = ud.UnixSockPath(os.Getenv("ENV_SOCK_PATH"))
var addr ud.StreamAddr = path2unixSock.ToAddr()

func sub(_ctx context.Context) error {
	listener, err := addr.Listen()
	if nil != err {
		return err
	}
	defer listener.Close()

	conn, err := listener.AcceptUnix()
	if nil != err {
		return err
	}
	defer conn.Close()

	var buf [65536]byte

	for {
		_, err := conn.Read(buf[:])
		if nil != err {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
	}
}

func main() {
	err := sub(context.Background())
	if nil != err {
		log.Printf("%v\n", err)
	}
}
