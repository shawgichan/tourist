package gapi

import (
	"fmt"

	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/pb"
	"github.com/shawgichan/tourist/token"
	"github.com/shawgichan/tourist/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedTouristServer
	Store           db.Store
	TokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {

	tokenMaker, err := token.NewPastoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil, fmt.Errorf("cannot create token %w", err)
	}
	server := &Server{
		Store:           store,
		TokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
