package rest

import (
	"context"
	demo_go_swagger_api "gateway"

	"gateway/gen/restapi/operations"
	"gateway/gen/restapi/operations/health_check"

	"gateway/internal/handlers"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *demo_go_swagger_api.Runtime, api *operations.GatewayAPI, apiHandler handlers.Handler) {
	rt.Logger().Info("Initialize Route")

	api.HealthCheckGetHealthCheckHandler = health_check.GetHealthCheckHandlerFunc(func(ghcp health_check.GetHealthCheckParams) middleware.Responder {

		result, err := apiHandler.GetHello(context.Background(), rt)

		if err != nil {
			return health_check.NewGetHealthCheckOK()
		}
		
		return health_check.NewGetHealthCheckOK().WithPayload(result)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}
}
