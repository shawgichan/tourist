package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/pb"
	"github.com/shawgichan/tourist/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userCheck, err := server.Store.CheckUsernameAndEmail(ctx, db.CheckUsernameAndEmailParams{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
	})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Errorf(codes.Internal, "error checking username and email: %s", err)
	}

	if userCheck.EmailPresent.Bool {
		return nil, status.Errorf(codes.AlreadyExists, "email already exists")
	}
	if userCheck.UsernamePresent.Bool {
		return nil, status.Errorf(codes.AlreadyExists, "username already exists")
	}

	profileArgs := db.CreateProfileParams{
		FirstName:       req.Name,
		LastName:        "",
		AddressesID:     0,
		ProfileImageUrl: "",
		PhoneNumber:     "",
		CompanyNumber:   "",
		WhatsappNumber:  "",
		Gender:          0,
		AllLanguagesID:  []int64{},
		RefNo:           utils.GenerateReferenceNumber("PROF_"),
		CoverImageUrl:   pgtype.Text{},
	}

	profile, err := server.Store.CreateProfile(ctx, profileArgs)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating profile: %s", err)
	}
	encryptedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password: %s", err)
	}
	password := string(encryptedPassword)

	userArgs := db.CreateUserParams{
		Email:          req.GetEmail(),
		Username:       req.GetUsername(),
		HashedPassword: password,
		Status:         0,
		RolesID:        1,
		ProfilesID:     profile.ID,
		UserTypesID:    0,
	}

	user, err := server.Store.CreateUser(ctx, userArgs)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating user: %s", err)
	}
	// resp := registerResponse{
	// 	Email:       user.Email,
	// 	Username:    user.Username,
	// 	Status:      user.Status,
	// 	RolesID:     pgtype.Int8{},
	// 	ProfilesID:  user.ProfilesID,
	// 	UserTypesID: user.UserTypesID,
	// }
	resp := &pb.CreateUserResponse{
		User: &pb.User{
			Username:          user.Username,
			Name:              profile.FirstName,
			Email:             user.Email,
			PasswordChangedAt: &timestamppb.Timestamp{},
			CreatedAt:         &timestamppb.Timestamp{},
		},
	}

	return resp, nil
}
