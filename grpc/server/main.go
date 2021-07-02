package main

import (
	"context"
	"fmt"
	"github.com/shiv3/opa-examples/http/server/grpc/grpcpb"
	"google.golang.org/grpc"
	"net"
)

type GrpcServiceImpl struct {
	grpcpb.UnimplementedGPRCServiceAPIServer
}

func (s GrpcServiceImpl)  Ping(ctx context.Context, in *grpcpb.PingRequest) (*grpcpb.PingResponse, error){
	return nil, nil
}
func (s GrpcServiceImpl)  ShowUser(ctx context.Context, in *grpcpb.ShowUserRequest) (*grpcpb.ShowUserResponse, error){
	return nil, nil
}
func (s GrpcServiceImpl)  ShowUsers(ctx context.Context, in *grpcpb.ShowUsersRequest) (*grpcpb.ShowUsersResponse, error){
	return nil, nil
}
func (s GrpcServiceImpl)  RegistUser(ctx context.Context, in *grpcpb.RegistUserRequest) (*grpcpb.RegistUserResponse, error){
	return nil, nil
}


func initial() error{
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	grpcpb.RegisterGPRCServiceAPIServer(grpcServer,GrpcServiceImpl{})

	return grpcServer.Serve(listener)

}

func main() {
	initial()
}
