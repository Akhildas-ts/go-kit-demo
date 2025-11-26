package task

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	spec "github.com/yourusername/task-service/spec/task"
)

func MakeEndpoints(bl *BL) spec.Endpoints {
	return spec.Endpoints{
		CreateTaskEndpoint:   MakeCreateTaskEndpoint(bl),
		ListTasksEndpoint:    MakeListTasksEndpoint(bl),
		GetTaskEndpoint:      MakeGetTaskEndpoint(bl),
		CompleteTaskEndpoint: MakeCompleteTaskEndpoint(bl),
	}
}

func MakeCreateTaskEndpoint(bl *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*spec.CreateTaskRequest)
		return bl.CreateTask(ctx, req)
	}
}

func MakeListTasksEndpoint(bl *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*spec.ListTasksRequest)
		return bl.ListTasks(ctx, req)
	}
}

func MakeGetTaskEndpoint(bl *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*spec.GetTaskRequest)
		return bl.GetTask(ctx, req)
	}
}

func MakeCompleteTaskEndpoint(bl *BL) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*spec.CompleteTaskRequest)
		return bl.CompleteTask(ctx, req)
	}
}
