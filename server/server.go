package server

import (
	"fmt"
	"net"

	"github.com/atmom/nodego/config"
	nodego "github.com/atmom/nodego/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// start server.
func Start(config *config.NodeServerConfig) {
	if config.Port < 1024 || config.Port > 65536 {
		logrus.Fatalf("Nodego server start error, invalid server port %d.", config.Port)
	}

	// listen tcp.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.Fatalf("Nodego server start error %v.", err)
	}

	// create server.
	s := grpc.NewServer()
	nodego.RegisterNodegoServer(s, &service{Config: config})

	logrus.Infof("Nodego server start success, listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		logrus.Fatalf("Nodego server start error %v.", err)
	}
}
