package main

import (
	"net/http"
	"os"

	gateway "gateway"
	"gateway/configs"
	"gateway/gen/restapi"
	"gateway/gen/restapi/operations"
	"gateway/internal/handlers"
	"gateway/internal/rest"

	"github.com/casualjim/middlewares"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/justinas/alice"
	"github.com/sirupsen/logrus"
)

var mainFlags = struct {
	AppConfig string `long:"config" description:"Main application configuration YAML path"`
}{}

func main() {
	log := initLogger()
	cfg := configs.Get()
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGatewayAPI(swaggerSpec)
	api.Logger = log.Infof
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "App Flags",
			LongDescription:  "",
			Options:          &mainFlags,
		},
	}

	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "demo go swagger api "
	parser.LongDescription = "demo go swagger api"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	rt, err := gateway.NewRuntime(cfg, log, false)
	if err != nil {
		log.Fatalln(err)
	}

	h := handlers.NewHandler()

	rest.Route(rt, api, h)

	handler := alice.New(
		middlewares.NewRecoveryMW("gateway", log),
		middlewares.NewProfiler,
	).Then(rest.SetupGlobalMiddleware(cfg, api.Serve(setupMiddlewares)))

	server.SetHandler(handler)
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func initLogger() logrus.FieldLogger {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	return logrus.StandardLogger()
}
