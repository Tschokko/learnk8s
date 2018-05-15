package controller

import (
	"log"
	"net"

	"github.com/tschokko/learnk8s/pkg/controller/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ServiceControllerServer struct{}

func NewServer() (*ServiceControllerServer, error) {
	s := &ServiceControllerServer{}
	return s, nil
}

func (s *ServiceControllerServer) RegisterService(ctx context.Context, in *api.RegisterServiceRequest) (*api.RegisterServiceResponse, error) {
	log.Printf("service with ID '%s' registered.", in.ServiceID)
	return &api.RegisterServiceResponse{Success: true}, nil
}

func (s *ServiceControllerServer) Run(l net.Listener) error {
	grpcServer := grpc.NewServer()
	api.RegisterServiceControllerServer(grpcServer, s)
	return grpcServer.Serve(l)
}
