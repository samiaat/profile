package main

import (
	"backend/pkg/api"
	"backend/pkg/api/post"
	"backend/pkg/db/sqlite"
	"backend/pkg/middleware"
	"log"
	"net/http"
)

func main() {
	db, err := sqlite.ConnectDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}
	defer sqlite.CloseDB()

	if err := sqlite.ApplyMigrations(db); err != nil {
		log.Fatalf("Could not apply migrations: %v\n", err)
	}

	appCore := middleware.NewAppCore(db)
	defer appCore.Close()

	go appCore.Hub.Run()

	mux := http.NewServeMux()



	// User routes
	mux.HandleFunc("/api/user/", api.GetUserHandler)
	mux.HandleFunc("/api/users", api.GetUsersHandler)
	mux.HandleFunc("/api/user/posts", post.GetUserPostsHandler)
	
	// Apply middlewares
	handler := middleware.CorsMiddleware(mux)
	rateLimiter := middleware.NewRateLimiter()
	handler = rateLimiter.RateLimitMiddleware(handler)
	handler = middleware.ErrorHandlerMiddleware(handler)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}