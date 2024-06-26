package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	graph_server "github.com/shawgichan/tourist/cmd/graph"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/token"
)

type Server struct {
	Store         db.Store
	StoreWithPool db.SQLStore
	TokenMaker    token.Maker
	Pool          *pgxpool.Pool
	Router        *gin.Engine
}

func NewServer(store db.Store, pool *pgxpool.Pool) (*Server, error) {

	tokenMaker, err := token.NewPastoMaker("11122233344455566677788899900012")
	if err != nil {
		return nil, fmt.Errorf("cannot create token %w", err)
	}
	router := gin.Default()
	server := &Server{Store: store, TokenMaker: tokenMaker, Pool: pool, Router: router}
	server.Router = router

	//router.MaxMultipartMemory = 100 << 2

	//rl := middleware.NewRateLimiter(10, 100)
	//router.Use(rl.Limiter())

	RouteServer := NewRouteServer(store, router, pool)
	userServer := NewUserServer(store, tokenMaker, router, pool)
	RouteServer.ApiRoutes()
	userServer.UserRoutes()
	graph := graph_server.NewGraphServer(server.Store, server.StoreWithPool, server.Router, server.Pool)
	graph.GraphServer()
	log.Println("testing total routes", len(server.Router.Routes()))

	//router.Use(middleware.TimeoutMiddleware())
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
