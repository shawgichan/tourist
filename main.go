package main

import (
	"log"
	"net"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	api "github.com/shawgichan/tourist/cmd"
	connect "github.com/shawgichan/tourist/db"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/gapi"
	"github.com/shawgichan/tourist/pb"
	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	gotenv.Load()
}

func main() {

	pool := connect.ConnectToDb()
	store := db.NewStore(pool)
	runGrpcServer(store)

}

func runGrpcServer(store db.Store) {
	server, err := gapi.NewServer(store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTouristServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}
	log.Println("Server started on port", os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

func runGinServer(store db.Store, pool *pgxpool.Pool) {
	server, err := api.NewServer(store, pool)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error while Starting server .")
	}
}
