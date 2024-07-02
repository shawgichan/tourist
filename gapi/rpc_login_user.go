package gapi

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shawgichan/tourist/pb"
	"github.com/shawgichan/tourist/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.Store.GetUserByName(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "username not found %s", err)
		} else {
			return nil, status.Errorf(codes.Internal, "error while getting user %s", err)
		}
	}

	passwordError := utils.CheckPassword(req.Password, user.HashedPassword)
	if passwordError != nil {
		return nil, status.Errorf(codes.PermissionDenied, "wrong password %s", err)
	}

	accessToken, err2 := server.TokenMaker.CreateToken(user.Username, 2*time.Hour)
	if err2 != nil {
		return nil, status.Errorf(codes.Internal, "error generating token %s", err2)
	}

	// rsp = loginResponse{
	// 	UserID:    user.ID,
	// 	Token:     accessToken,
	// 	Email:     user.Email,
	// 	UserName:  user.Username,
	// 	FirstName: "",
	// 	LastName:  "",
	// }
	rsp := &pb.LoginUserResponse{
		User: &pb.User{
			Username:          user.Username,
			Name:              "",
			Email:             user.Email,
			PasswordChangedAt: &timestamppb.Timestamp{},
			CreatedAt:         &timestamppb.Timestamp{},
		},
		SessionId:          "",
		AccessToken:        accessToken,
		AccessTokenExpires: &timestamppb.Timestamp{},
	}

	return rsp, nil
}
