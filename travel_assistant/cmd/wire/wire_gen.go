// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/Wenrh2004/travel_assistant/internal/adapter"
	"github.com/Wenrh2004/travel_assistant/internal/adapter/server"
	"github.com/Wenrh2004/travel_assistant/internal/domain/agent"
	"github.com/Wenrh2004/travel_assistant/internal/infrastructure/llm"
	"github.com/Wenrh2004/travel_assistant/internal/infrastructure/third"
	"github.com/Wenrh2004/travel_assistant/internal/infrastructure/third/tool"
	"github.com/Wenrh2004/travel_assistant/pkg/application"
	"github.com/Wenrh2004/travel_assistant/pkg/application/server/http"
	"github.com/Wenrh2004/travel_assistant/pkg/util/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*application.App, func(), error) {
	chatModelService := llm.NewChatModelService(viperViper)
	client := third.NewAmapClient(viperViper)
	poiService := tool.NewPOIService(client)
	routerService := tool.NewRouterService(client)
	weatherService := tool.NewWeatherService(client)
	chatService := agent.NewDomain(viperViper, chatModelService, poiService, routerService, weatherService)
	llmHandler := adapter.NewLLMHandler(logger, chatService)
	httpServer := server.NewHTTPServer(logger, llmHandler)
	app := newApp(httpServer, viperViper)
	return app, func() {
	}, nil
}

// wire.go:

var infrastructureSet = wire.NewSet(third.NewAmapClient, tool.NewPOIService, tool.NewRouterService, tool.NewWeatherService, llm.NewChatModelService)

var domainSet = wire.NewSet(agent.NewDomain)

var adapterSet = wire.NewSet(adapter.NewLLMHandler)

var serverSet = wire.NewSet(server.NewHTTPServer)

// build App
func newApp(
	httpServer *http.Server,
	conf *viper.Viper,
) *application.App {
	return application.NewApp(application.WithServer(httpServer), application.WithName(conf.GetString("app.name")))
}
