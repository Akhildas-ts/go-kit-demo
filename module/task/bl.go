package task

import (
	"context"
	"errors"

	"github.com/google/uuid"
	spec "github.com/yourusername/task-service/spec/task"
)

type BL struct {
	repo *Repository
}

func NewBL(repo *Repository) *BL {
	return &BL{repo: repo}
}

func (b *BL) CreateTask(ctx context.Context, req *spec.CreateTaskRequest) (*spec.CreateTaskResponse, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}

	id, err := b.repo.CreateTask(ctx, req.Title, req.Description)
	if err != nil {
		return nil, err
	}

	return &spec.CreateTaskResponse{
		ID:      id,
		Message: "Task created successfully",
	}, nil
}

func (b *BL) ListTasks(ctx context.Context, req *spec.ListTasksRequest) (*spec.ListTasksResponse, error) {
	tasks, err := b.repo.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	return &spec.ListTasksResponse{Tasks: tasks}, nil
}

func (b *BL) GetTask(ctx context.Context, req *spec.GetTaskRequest) (*spec.GetTaskResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	task, err := b.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &spec.GetTaskResponse{Task: *task}, nil
}

func (b *BL) CompleteTask(ctx context.Context, req *spec.CompleteTaskRequest) (*spec.CompleteTaskResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	err = b.repo.CompleteTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return &spec.CompleteTaskResponse{
		Message: "Task marked as complete",
	}, nil
}
