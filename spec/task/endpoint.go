package task

import "github.com/go-kit/kit/endpoint"

type Endpoints struct {
	CreateTaskEndpoint   endpoint.Endpoint
	ListTasksEndpoint    endpoint.Endpoint
	GetTaskEndpoint      endpoint.Endpoint
	CompleteTaskEndpoint endpoint.Endpoint
}
