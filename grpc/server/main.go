package main

//go:generate sqlite3 schema/models.sqlite < schema/ddl.sql
//go:generate sqlboiler sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shiv3/opa-examples/http/server/grpc/grpcpb"
	"github.com/shiv3/opa-examples/http/server/models"
	"github.com/volatiletech/sqlboiler/boil"
	"google.golang.org/grpc"
	"net"
)

type GrpcServiceImpl struct {
	db *sql.DB
	grpcpb.UnimplementedGPRCServiceAPIServer
}

func (s GrpcServiceImpl) Ping(ctx context.Context, in *grpcpb.PingRequest) (*grpcpb.PingResponse, error) {
	return &grpcpb.PingResponse{
		Message: "ok",
	}, nil
}
func (s GrpcServiceImpl) ShowUser(ctx context.Context, in *grpcpb.ShowUserRequest) (*grpcpb.ShowUserResponse, error) {
	user, err := models.FindUser(ctx, s.db, in.GetId())
	if err != nil {
		return nil, err
	}
	return &grpcpb.ShowUserResponse{
		User: &grpcpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
func (s GrpcServiceImpl) ShowUsers(ctx context.Context, in *grpcpb.ShowUsersRequest) (*grpcpb.ShowUsersResponse, error) {
	users, err := models.Users().All(context.Background(), s.db)
	if err != nil {
		return nil, err
	}
	res := make([]*grpcpb.User, 0, len(users))
	for _, user := range users {
		res = append(res, &grpcpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &grpcpb.ShowUsersResponse{Users: res}, nil
}
func (s GrpcServiceImpl) RegistUser(ctx context.Context, in *grpcpb.RegistUserRequest) (*grpcpb.RegistUserResponse, error) {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	u := models.User{
		ID:    in.GetUser().GetId(),
		Name:  in.GetUser().GetName(),
		Email: in.GetUser().GetEmail(),
	}

	if err := u.Insert(context.Background(), tx, boil.Infer()); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &grpcpb.RegistUserResponse{}, nil
}

func initial() error {
	db, err := sql.Open("sqlite3", "schema/models.sqlite")
	if err != nil {
		return err
	}

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	grpcpb.RegisterGPRCServiceAPIServer(grpcServer, GrpcServiceImpl{db: db})

	return grpcServer.Serve(listener)

}

func main() {
	err := initial()
	if err != nil {
		panic(err)
	}
}
