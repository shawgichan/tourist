package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/repo/places"
)

type RouteServer struct {
	Store  db.Store
	Pool   *pgxpool.Pool
	Router *gin.Engine
}

func NewRouteServer(store db.Store, router *gin.Engine, pool *pgxpool.Pool) *RouteServer {
	return &RouteServer{
		Store:  store,
		Pool:   pool,
		Router: router,
	}
}

func (server *RouteServer) ApiRoutes() {
	places := places.NewPlaceServer(server.Store)

	d := server.Router.Group("api/places")
	{
		d.POST("/createNewPlace", places.CreateNewPlace)
		d.GET("/getAllPlaces", places.GetAllPlacesPlace)
	}
}
