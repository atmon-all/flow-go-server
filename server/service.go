package server

import (
	"context"

	"github.com/atmom/nodego/config"
	nodego "github.com/atmom/nodego/grpc"
	"github.com/sirupsen/logrus"
)

type service struct {
	nodego.UnimplementedNodegoServer
	Config *config.NodeServerConfig
}

// Execute next flow node.
func (service *service) GetData(ctx context.Context, request *nodego.ExecuteNextRequest) (*nodego.ExecuteNextResponse, error) {
	logrus.Infof("Received: %v", request)
	return &nodego.ExecuteNextResponse{Code: 0, Message: "success"}, nil
}
