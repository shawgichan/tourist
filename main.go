package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	api "github.com/shawgichan/tourist/cmd"
	connect "github.com/shawgichan/tourist/db"
	db "github.com/shawgichan/tourist/db/sqlc"
	_ "github.com/shawgichan/tourist/doc/statik"
	"github.com/shawgichan/tourist/gapi"
	"github.com/shawgichan/tourist/pb"
	"github.com/shawgichan/tourist/worker"
	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	gotenv.Load()
}

func main() {

	pool := connect.ConnectToDb()
	store := db.NewStore(pool)
	// redisOpt := asynq.RedisClientOpt{
	// 	Addr: os.Getenv("REDIS_ADDR"),
	// }

	//taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	//go runTaskProcessor(redisOpt, store)
	//go runGatewayServer(store, taskDistributor)
	runGinServer(store, pool)
	//runGrpcServer(store, taskDistributor)

}

func runGrpcServer(store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(store, taskDistributor)
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

func runGatewayServer(store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(store, taskDistributor)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterTouristHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("cannot create statik fs : ", err)
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))

	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}
	log.Println("HttpServer started on port", os.Getenv("PORT"))
	if err := http.Serve(listener, mux); err != nil {
		log.Fatal("Cannot start HTTP server:", err)
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

func runTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) {
	processor := worker.NewRedisTaskProcessor(redisOpt, store)
	log.Println("start task processor")
	err := processor.Start()
	if err != nil {
		log.Fatal("cannot start task processor: ", err)
	}
}
