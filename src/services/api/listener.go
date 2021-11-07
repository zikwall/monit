package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/zikwall/monit/src/pkg/logger"
)

const (
	ListenerTCP = iota + 1
	ListenerUDS
)

func listenToUnix(bind string) (net.Listener, error) {
	_, err := os.Stat(bind)

	if err == nil {
		err = os.Remove(bind)

		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	return net.Listen("unix", bind)
}

func maybeChmodSocket(c context.Context, sock string) {
	ctx, cancel := context.WithTimeout(c, 500*time.Millisecond)
	defer cancel()

	// on Linux and similar systems, there may be problems with the rights to the UDS socket
	go func() {
		var tryCount uint

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Millisecond * 100):
				if err := os.Chmod(sock, 0o666); err == nil {
					logger.Warning(err)
				} else {
					_, err := os.Stat(sock)
					// if the file exists and it already has permissions
					if err == nil {
						return
					}
				}

				tryCount++
				if tryCount > 5 {
					return
				}
			}
		}
	}()

	_ = os.Chmod(sock, 0666)
}
