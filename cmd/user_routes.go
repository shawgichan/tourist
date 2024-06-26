package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/repo/user"
	"github.com/shawgichan/tourist/token"
)

type UserServer struct {
	Store      db.Store
	Pool       *pgxpool.Pool
	TokenMaker token.Maker
	Router     *gin.Engine
}

func NewUserServer(store db.Store, token token.Maker, router *gin.Engine, pool *pgxpool.Pool) *UserServer {
	return &UserServer{
		Store:      store,
		Pool:       pool,
		TokenMaker: token,
		Router:     router,
	}
}
func (server *UserServer) UserRoutes() {
	users := user.NewUserServer(server.Store)

	//auth := middleware.AuthMiddleware(server.TokenMaker)
	d := server.Router.Group("api/auth")
	{
		d.POST("/registerUser", users.RegisterUser)
		d.GET("/loginUser", users.LoginUser)
	}
}
