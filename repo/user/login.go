package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/utils"
)

type loginUserRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type loginResponse struct {
	UserID    int64
	Token     string
	Email     string
	UserName  string
	FirstName string
	LastName  string
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req loginUserRequest
	var user db.User
	var rsp loginResponse

	if err := ctx.ShouldBind(&req); err != nil {
		er := utils.BindingFormError(err, req)
		log.Println("testing er", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": er})
		return
	}
	user, err := server.SQLStore.GetUserByName(ctx, req.UserName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errors.New("not found"))
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err})
		}
	}

	// passwordError := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword.String), []byte(req.Password))
	// if passwordError != nil {
	// 	ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	// 	return
	// }
	passwordError := utils.CheckPassword(req.Password, user.HashedPassword.String)
	if passwordError != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(passwordError))
		return
	}

	// accessToken, _, err2 := server.TokenMaker.CreateToken(user.Username, 2*time.Hour)
	// if err2 != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	rsp = loginResponse{
		UserID: user.ID,
		//Token:     accessToken,
		Email:     user.Email,
		UserName:  user.Username,
		FirstName: "",
		LastName:  "",
	}
	ctx.JSON(http.StatusOK, rsp)
}
