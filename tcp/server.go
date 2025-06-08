package tcp

import (
	"context"
	"go-Redis/interface/tcp"
	"go-Redis/lib/logger"
	"net"
)

type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	closeChan := make(chan struct{})
	listener, err := net.Listen("tcp", cfg.Address)

	if err != nil {
		return err
	}

	logger.Info("Start listen")

	ListenAndServe(listener, handler, closeChan)
	return nil
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {
	ctx := context.Background()
	for true {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		logger.Info("Accepted link")
		go func() {
			handler.Handle(ctx, conn)
		}()
	}
}
