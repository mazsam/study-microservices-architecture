package handlers

import (
	"context"
	demo_go_swagger_api "gateway"
	"gateway/gen/models"
)

// GetHello implements Handler
func (h *handler) GetHello(ctx context.Context, rt *demo_go_swagger_api.Runtime) (result []*models.Hello, err error) {
	for _, v := range rt.Cfg.ErrorCode {
		result = append(result, &models.Hello{
			Name:  v.Message,
			Title: v.Message,
		})
	}

	return result, nil
}
