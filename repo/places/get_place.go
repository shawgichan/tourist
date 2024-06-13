package places

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllPlacesPlace(ctx *gin.Context) {

	place, err := server.SQLStore.GetPlaces(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, place)
}
