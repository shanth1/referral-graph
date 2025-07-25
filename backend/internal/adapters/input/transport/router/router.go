package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shanth1/graph/internal/core/domain"
	"github.com/shanth1/graph/internal/core/ports"
)

func NewRouter(service ports.ReferralService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/graph", getGraphHandler(service))
	mux.HandleFunc("/api/health", healthCheckHandler)

	handler := corsMiddleware(loggingMiddleware(mux))

	return handler
}

func getGraphHandler(service ports.ReferralService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nodes, edges, err := service.GetGraphData()
		if err != nil {
			http.Error(w, "Failed to get graph data", http.StatusInternalServerError)
			return
		}

		response := struct {
			Nodes []domain.Node `json:"nodes"`
			Edges []domain.Edge `json:"edges"`
		}{
			Nodes: nodes,
			Edges: edges,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Middlewares
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
