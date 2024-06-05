package grpcserver

import (
	"context"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {

	dbUser, err := server.uc.CreateUser(ctx, req.GetUsername(), req.GetEmail(), req.GetFullname(), req.GetPassword(), req.GetRole())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %v", err)
	}
	return &pb.User{
		ID:        dbUser.ID.String(),
		Username:  dbUser.Username,
		Email:     dbUser.Email.String,
		Fullname:  dbUser.DisplayName.String,
		Password:  dbUser.Password,
		Role:      dbUser.Role,
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
		Id:          user.ID.String(),
		Username:    req.GetUsername(),
		Email:       user.Email.String,
		DisplayName: user.DisplayName.String,
		Role:        user.Role,
	}
	maker, _ := paseto.NewPasetoMaker()
	token, _, err := maker.CreateToken(user.Username, time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't login with user name and password: %v", err)

	}
	rsp.Token = token
	return rsp, nil
}

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserIdRequest) (*pb.User, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "User Id is invalid: %v", err)
	}
	user, err := server.uc.GetUser(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "User doesn't match: %v", err)
	}
	return &pb.User{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email.String,
		Fullname:  user.DisplayName.String,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: timestamppb.New(user.Created),
		// Add more fields as needed
	}, nil
}
