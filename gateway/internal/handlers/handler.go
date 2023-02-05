package handlers

import (
	"context"
	"gateway/gen/models"

	demo_go_swagger_api "gateway"
)

type Handler interface {
	HelloHandler
}

type HelloHandler interface {
	GetHello(ctx context.Context, rt *demo_go_swagger_api.Runtime) ([]*models.Hello, error)
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}
