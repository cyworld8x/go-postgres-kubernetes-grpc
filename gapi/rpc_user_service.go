package gapi

import (
	"context"
	"time"

	entity "github.com/cyworld8x/go-postgres-kubernetes-grpc/db/sqlc"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/paseto"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"
	password "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	pwd, err := password.HashPassword(req.GetPassword())
	if err != nil {
		return &pb.User{}, err
	}

	arg := entity.CreateUserParams{
		Username: pgtype.Text{String: req.GetUsername(), Valid: true},
		Email:    pgtype.Text{String: req.GetEmail(), Valid: true},
		Fullname: pgtype.Text{String: req.GetFullname(), Valid: true},
		Password: pgtype.Text{String: pwd, Valid: true},
		Role:     pgtype.Text{String: req.GetRole(), Valid: true},
	}

	dbUser, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %v", err)
	}
	return &pb.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username.String,
		Email:     dbUser.Email.String,
		Fullname:  dbUser.Fullname.String,
		Password:  dbUser.Password.String,
		Role:      dbUser.Role.String,
		CreatedAt: timestamppb.New(dbUser.CreatedAt.Time),
		// Add more fields as needed
	}, nil
}

func (server *Server) GetLogin(ctx context.Context, req *pb.GetLoginRequest) (*pb.GetLoginResponse, error) {
	usrName := pgtype.Text{String: req.GetUsername(), Valid: true}
	user, err := server.store.GetLogin(ctx, usrName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "User name doesn't match: %v", err)

	}

	err = password.CheckPassword(req.GetPassword(), user.Password.String)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't login with user name and password: %v", err)
	}

	rsp := &pb.GetLoginResponse{
		Id:       user.ID,
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	maker, _ := paseto.NewPasetoMaker()
	token, _, err := maker.CreateToken(user.Username.String, time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't login with user name and password: %v", err)

	}
	rsp.Token = token
	return rsp, nil
}

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserIdRequest) (*pb.User, error) {
	user, err := server.store.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get user: %v", err)
	}
	return &pb.User{
		ID:        user.ID,
		Username:  user.Username.String,
		Email:     user.Email.String,
		Fullname:  user.Fullname.String,
		Password:  user.Password.String,
		Role:      user.Role.String,
		CreatedAt: timestamppb.New(user.CreatedAt.Time),
		// Add more fields as needed
	}, nil
}
