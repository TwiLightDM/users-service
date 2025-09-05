package grpc

import (
	"fmt"
	userpb "github.com/TwiLightDM/project-protos/proto/user"
	"github.com/TwiLightDM/users-service/internal/user"
	"google.golang.org/grpc"
	"net"
)

func RunGRPC(svc user.Service) error {
	// 1. net.Listen на ":50051"
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	// 2. grpc.NewServer()
	grpcSrv := grpc.NewServer()

	// 3. userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	// 4. grpcSrv.Serve(listener)
	fmt.Println("gRPC server is running on :50051")
	if err = grpcSrv.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
