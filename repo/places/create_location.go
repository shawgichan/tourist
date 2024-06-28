package places

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/utils"
)

type createLocationRequest struct {
	Lat string `json:"lat" binding:"required"`
	Lon string `json:"lon" binding:"required"`
}

func (server *Server) CreateLocation(ctx *gin.Context) {
	var req createLocationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	output, err := server.SQLStore.CreateLocation(ctx, db.CreateLocationParams{
		Lat:       req.Lat,
		Lng:       req.Lon,
		CreatedAt: time.Now().Round(time.Second),
		UpdatedAt: time.Now().Round(time.Second),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(output))
}
