package graph_server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/graph/places"
)

type GraphServer struct {
	Store    db.Store
	SQLStore *db.SQLStore
	//TokenMaker token.Maker
	Router *gin.Engine
	Conn   *pgxpool.Pool
}

func NewGraphServer(store db.Store, SQLStore db.SQLStore, router *gin.Engine, conn *pgxpool.Pool) *GraphServer {
	return &GraphServer{
		Store:    store,
		SQLStore: &SQLStore,
		//TokenMaker: token,
		Router: router,
		Conn:   conn,
	}
}

// ! *************************  DO NOT EDIT:  For Graphql server  ***************************
func (server *GraphServer) GraphServer() {

	// server.Router.Use(middleware.AuthMiddlewareForGraph(server.TokenMaker))
	//auth := middleware.AuthMiddlewareForGraph(server.TokenMaker)

	r := server.Router.Group("api/")
	{
		r.GET("api", PlaygroundHandler())
		r.POST("places", server.GraphqlPlacesHandler())

	}

}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (server *GraphServer) GraphqlPlacesHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(
		places.NewExecutableSchema(
			places.Config{
				Resolvers: &places.Resolver{
					Store: server.Store,
				},
				Directives: places.DirectiveRoot{},
				Complexity: places.ComplexityRoot{},
			},
		))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
