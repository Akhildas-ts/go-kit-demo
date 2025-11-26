package task

import "github.com/google/uuid"

// CREATE TASK
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}

// LIST TASKS
type ListTasksRequest struct{}

type ListTasksResponse struct {
	Tasks []Task `json:"tasks"`
}

// GET TASK BY ID
type GetTaskRequest struct {
	ID string `json:"id"`
}

type GetTaskResponse struct {
	Task Task `json:"task"`
}

// COMPLETE TASK
type CompleteTaskRequest struct {
	ID string `json:"id"`
}

type CompleteTaskResponse struct {
	Message string `json:"message"`
}

// MODELS
type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   string    `json:"created_at"`
}
