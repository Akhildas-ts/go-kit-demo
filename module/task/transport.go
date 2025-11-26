package task

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httptransport "github.com/go-kit/kit/transport/http"

	spec "github.com/yourusername/task-service/spec/task"
)

func MakeHTTPHandler(bl *BL) http.Handler {
	endpoints := MakeEndpoints(bl)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CREATE TASK
	r.Post("/v1/tasks", httptransport.NewServer(
		endpoints.CreateTaskEndpoint,
		decodeCreateTaskRequest,
		encodeJSONResponse,
	).ServeHTTP)

	// LIST TASKS
	r.Get("/v1/tasks", httptransport.NewServer(
		endpoints.ListTasksEndpoint,
		decodeListTasksRequest,
		encodeJSONResponse,
	).ServeHTTP)

	// GET TASK BY ID
	r.Get("/v1/tasks/{id}", httptransport.NewServer(
		endpoints.GetTaskEndpoint,
		decodeGetTaskRequest,
		encodeJSONResponse,
	).ServeHTTP)

	// COMPLETE TASK
	r.Put("/v1/tasks/{id}/complete", httptransport.NewServer(
		endpoints.CompleteTaskEndpoint,
		decodeCompleteTaskRequest,
		encodeJSONResponse,
	).ServeHTTP)

	return r
}

// DECODERS
func decodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req spec.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func decodeListTasksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return &spec.ListTasksRequest{}, nil
}

func decodeGetTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return &spec.GetTaskRequest{ID: id}, nil
}

func decodeCompleteTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")
	return &spec.CompleteTaskRequest{ID: id}, nil
}

// ENCODER
func encodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
