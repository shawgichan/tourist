package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/repo/utils"
)

type Server struct {
	SQLStore db.Store
}

func NewPlaceServer(store db.Store) *Server {
	return &Server{SQLStore: store}
}

type registerUserRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Email    string `form:"email" binding:"required"`
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
		ctx.JSON(http.StatusBadRequest, errors.New("email already existing"))
		return
	}
	if userCheck.UsernamePresent.Bool {
		ctx.JSON(http.StatusBadRequest, errors.New("username already existing"))
		return
	}

	ctx.JSON(http.StatusOK, "user created successfully")

}
