package server

import (
	"fmt"
	"go-inertia/internal/database"
	"go-inertia/internal/middleware"
	"net/http"
	"os"
	"strconv"
	"time"

	inertiaInit "go-inertia/internal/inertia"

	_ "github.com/joho/godotenv/autoload"
	inertia "github.com/romsar/gonertia/v2"
)

type Server struct {
	port    int
	inertia *inertia.Inertia
	db      database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:    port,
		inertia: inertiaInit.InitInertia(),
		db:      database.New(),
	}

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Cors,
	)

	router := NewServer.RegisterRoutes()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      stack(router),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
