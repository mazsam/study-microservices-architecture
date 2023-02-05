// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gateway/gen/restapi/operations"
	"gateway/gen/restapi/operations/auth"
	"gateway/gen/restapi/operations/health_check"
	"gateway/gen/restapi/operations/user"
)

//go:generate swagger generate server --target ../../gen --name Gateway --spec ../../api/swagger.yml --principal models.Principal --exclude-main

func configureFlags(api *operations.GatewayAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GatewayAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.UserCreateProfileHandler == nil {
		api.UserCreateProfileHandler = user.CreateProfileHandlerFunc(func(params user.CreateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateProfile has not yet been implemented")
		})
	}
	if api.HealthCheckGetHealthCheckHandler == nil {
		api.HealthCheckGetHealthCheckHandler = health_check.GetHealthCheckHandlerFunc(func(params health_check.GetHealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation health_check.GetHealthCheck has not yet been implemented")
		})
	}
	if api.AuthLoginUserHandler == nil {
		api.AuthLoginUserHandler = auth.LoginUserHandlerFunc(func(params auth.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginUser has not yet been implemented")
		})
	}
	if api.AuthRegisterUserHandler == nil {
		api.AuthRegisterUserHandler = auth.RegisterUserHandlerFunc(func(params auth.RegisterUserParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.RegisterUser has not yet been implemented")
		})
	}
	if api.UserUpdateProfileHandler == nil {
		api.UserUpdateProfileHandler = user.UpdateProfileHandlerFunc(func(params user.UpdateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateProfile has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
