package main

import (
	"context"
	"grpcgqlgo/generated/graph"
	graphql "grpcgqlgo/graph/resolver"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	// config router options
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT},
		},
	))

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graphql.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Println(err)
		return gqlerror.Errorf("Internal server error!")
	})

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	e.GET("/*", echo.WrapHandler(srv))
	e.POST("/*", echo.WrapHandler(srv))

	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
