package grpc

import (
	"fmt"
	userpb "github.com/TwiLightDM/project-protos/proto/user"
	"github.com/TwiLightDM/users-service/internal/user"
	"google.golang.org/grpc"
	"net"
)

func RunGRPC(svc user.Service) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcSrv := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	fmt.Println("gRPC server is running on :50051")

	if err = grpcSrv.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
