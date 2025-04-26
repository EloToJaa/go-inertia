package server

import (
	"encoding/json"
	"go-inertia/internal/middleware"
	"log"
	"net/http"

	inertia "github.com/romsar/gonertia/v2"
)

func (s *Server) RegisterInertiaRoutes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /", s.HomeHandler)

	// Wrap the mux with CORS middleware
	return s.inertia.Middleware(router)
}

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()

	// Register routes
	router.Handle("GET /build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./frontend/public/build"))))
	router.HandleFunc("GET /health", s.healthHandler)

	router.HandleFunc("GET /hello", s.HelloWorldHandler)

	return router
}

func handleServerErr(w http.ResponseWriter, err error) {
	log.Printf("http error: %s\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("server error"))
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.inertia.Render(w, r, "home/index", inertia.Props{
		"text": "Inertia.js with Svelte and Go! 💙",
	})
	if err != nil {
		handleServerErr(w, err)
		return
	}
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.Health())
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
