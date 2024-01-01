package main

import (
	"bk/bk_api_gateway/config"
	"bk/bk_api_gateway/graph/generated"
	graph "bk/bk_api_gateway/graph/resolver"
	"bk/bk_api_gateway/pkg/logger"
	"bk/bk_api_gateway/services"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	fmt.Println(`H`,cfg.OrderServiceHost, `P`, cfg.OrderServicePort, `PORT:`, cfg.HTTPPort)
	grpcSvcs, err := services.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("bk_api_gateway", loggerLevel)
	defer logger.Cleanup(log)

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.Use(customCORSMiddleware())
	r.Use(GinContextToContextMiddleware())
	r.POST("/query", graphqlHandler(grpcSvcs, log))

	r.GET("/", playgroundHandler())

	r.Run(cfg.HTTPPort)
}

// Defining the Graphql handler
func graphqlHandler(svcs services.ServiceManagerI, log logger.LoggerI) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ServiceManagerI: svcs,
		LoggerI:         log,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
