package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/token"
	"github.com/shawgichan/tourist/utils"
)

type Server struct {
	SQLStore   db.Store
	Pool       *pgxpool.Pool
	TokenMaker token.Maker
}

func NewUserServer(store db.Store) *Server {
	return &Server{SQLStore: store}
}

type registerUserRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type registerResponse struct {
	Email       string      `json:"email"`
	Username    string      `json:"username"`
	Status      int64       `json:"status"`
	RolesID     pgtype.Int8 `json:"roles_id"`
	ProfilesID  int64       `json:"profiles_id"`
	UserTypesID int64       `json:"user_types_id"`
}

func (server *Server) RegisterUser(ctx *gin.Context) {
	var req registerUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		er := utils.BindingFormError(err, req)
		log.Println("testing er", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": er})
		return
	}

	//check email and username existance
	userCheck, err := server.SQLStore.CheckUsernameAndEmail(ctx, db.CheckUsernameAndEmailParams{
		Username: req.UserName,
		Email:    req.Email,
	})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if userCheck.EmailPresent.Bool {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errors.New("email already existing")))
		return
	}
	if userCheck.UsernamePresent.Bool {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(errors.New("username already existing")))
		return
	}

	profileArgs := db.CreateProfileParams{
		FirstName:       "",
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

	profile, err := server.SQLStore.CreateProfile(ctx, profileArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	encryptedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	password := string(encryptedPassword)

	userArgs := db.CreateUserParams{
		Email:          req.Email,
		Username:       req.UserName,
		HashedPassword: pgtype.Text{String: password, Valid: true},
		Status:         0,
		RolesID:        pgtype.Int8{},
		ProfilesID:     profile.ID,
		UserTypesID:    0,
	}

	user, err := server.SQLStore.CreateUser(ctx, userArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	resp := registerResponse{
		Email:       user.Email,
		Username:    user.Username,
		Status:      user.Status,
		RolesID:     pgtype.Int8{},
		ProfilesID:  user.ProfilesID,
		UserTypesID: user.UserTypesID,
	}

	ctx.JSON(http.StatusOK, resp)

}
