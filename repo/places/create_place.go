package places

import (
	"github.com/gin-gonic/gin"
	db "github.com/shawgichan/tourist/db/sqlc"
)

type Server struct {
	SQLStore db.Store
}

func NewPlaceServer(store db.Store) *Server {
	return &Server{SQLStore: store}
}

func (server *Server) CreateNewPlace(ctx *gin.Context) {
	var params db.CreatePlaceParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	place, err := server.SQLStore.CreatePlace(ctx, params)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, place)
}
