package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/shawgichan/tourist/db/sqlc"
)

type Server struct {
	Store         db.Store
	StoreWithPool db.SQLStore
	Pool          *pgxpool.Pool

	Router *gin.Engine
}

func NewServer(store db.Store, pool *pgxpool.Pool) (*Server, error) {

	router := gin.Default()
	server := &Server{Store: store, Pool: pool, Router: router}
	server.Router = router

	router.MaxMultipartMemory = 100 << 2

	log.Println("testing total routes", len(server.Router.Routes()))

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
