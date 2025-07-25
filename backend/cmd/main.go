package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/shanth1/graph/internal/adapters/input/transport/router"
	"github.com/shanth1/graph/internal/adapters/output/graphdb"
	"github.com/shanth1/graph/internal/common"
	"github.com/shanth1/graph/internal/config"
	"github.com/shanth1/graph/internal/core/ports"
	"github.com/shanth1/graph/internal/core/service"
)

func main() {
	// TODO: added to gotool
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// --- Сonfig ---
	cfg := config.MustGetConfig()

	logger := common.GetLogger()
	ctx = logger.WithContext(ctx)

	var graphRepo ports.GraphRepository
	if cfg.UseMocks {
		log.Println("Using MOCK repository")
		graphRepo = graphdb.NewMockRepository()
	} else {
		log.Println("Using Neo4j repository")
		repo, err := graphdb.NewNeo4jRepository(cfg.Neo4jURI, cfg.Neo4jUser, cfg.Neo4jPassword)
		if err != nil {
			log.Fatalf("cannot connect to neo4j: %v", err)
		}
		graphRepo = repo
	}

	referralService := service.NewReferralService(graphRepo)

	httpRouter := router.NewRouter(referralService)

	// TODO: move to router
	go func() {
		log.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", httpRouter); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-ctx.Done()

	logger.Info().Msg("shutting down server...")
}
