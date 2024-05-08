package grpcserver

import (
	"context"
	"strconv"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	pwd, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return &pb.User{}, err
	}

	dbUser, err := server.uc.CreateUser(ctx, req.GetUsername(), req.GetEmail(), req.GetFullname(), pwd, 0)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %v", err)
	}
	return &pb.User{
		Username:  dbUser.Username,
		Email:     dbUser.Email.String,
		Fullname:  dbUser.DisplayName.String,
		Password:  dbUser.Password,
		Role:      strconv.Itoa(int(dbUser.Role)),
		CreatedAt: timestamppb.New(dbUser.Created),
		// Add more fields as needed
	}, nil
}

func (server *Server) GetLogin(ctx context.Context, req *pb.GetLoginRequest) (*pb.GetLoginResponse, error) {
	user, err := server.uc.GetLogin(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "User name doesn't match: %v", err)

	}

	err = utils.CheckPassword(req.GetPassword(), user.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't login with user name and password: %v", err)
	}

	rsp := &pb.GetLoginResponse{
		Id:       0,
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	maker, _ := paseto.NewPasetoMaker()
	token, _, err := maker.CreateToken(user.Username, time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't login with user name and password: %v", err)

	}
	rsp.Token = token
	return rsp, nil
}
