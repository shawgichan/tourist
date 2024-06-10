package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	api "github.com/shawgichan/tourist/cmd"
	connect "github.com/shawgichan/tourist/db"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	pool := connect.ConnectToDb()
	store := db.NewStore(pool)
	server, err := api.NewServer(store, pool)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error while Starting server .")
	}

}
