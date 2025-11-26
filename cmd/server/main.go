package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	taskModule "github.com/yourusername/task-service/module/task"
	"github.com/yourusername/task-service/pkg/database"
)

func main() {
	// Connect to database
	db, err := database.Connect("postgres://postgres:akhil%40123@localhost/taskdb?sslmode=disable")

	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	// Create repository
	repo := taskModule.NewRepository(db)

	// Create business logic
	bl := taskModule.NewBL(repo)

	// Create HTTP handler
	handler := taskModule.MakeHTTPHandler(bl)

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
