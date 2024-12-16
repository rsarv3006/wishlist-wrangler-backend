package coresharedrestservice

import (
	"fmt"
	"log"
	"net"
	"net/http"
	coresharedconfig "wishlist-wrangler-api/src/core-shared/core-shared-config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

func SetupRouter(routes []Route) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			r.Get(route.Pattern, route.Handler)
		case http.MethodPost:
			r.Post(route.Pattern, route.Handler)
		case http.MethodPut:
			r.Put(route.Pattern, route.Handler)
		case http.MethodDelete:
			r.Delete(route.Pattern, route.Handler)
		}
	}

	return r
}

func StartServer(router *chi.Mux, startingPort int, portChan chan<- int) {
	portStr := fmt.Sprintf(":%d", startingPort)

	ln, err := net.Listen("tcp", portStr)
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}

	port := ln.Addr().(*net.TCPAddr).Port
	portChan <- port

	log.Printf("Starting server on port %d", port)
	if err := http.Serve(ln, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func StartServiceWithRoutes(config *coresharedconfig.ConfigDto, appName string, additionalRoutes []Route) {
	portChan := make(chan int, 1)
	defaultRoutes := []Route{
		{Method: http.MethodGet, Pattern: "/health", Handler: HealthCheck},
		{Method: http.MethodGet, Pattern: "/status", Handler: StatusCheck},
		{Method: http.MethodGet, Pattern: "/", Handler: HomeHandler},
		{Method: http.MethodGet, Pattern: "/swagger-ui/*", Handler: httpSwagger.Handler(
			httpSwagger.URL("/v3/api-docs"),
		)},
		{Method: http.MethodGet, Pattern: "/v3/api-docs", Handler: buildServerSwaggerHandler()},
	}

	routes := append(defaultRoutes, additionalRoutes...)
	router := SetupRouter(routes)

	go StartServer(router, config.ServicePort, portChan)

	select {}
}

func buildServerSwaggerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	}
}
