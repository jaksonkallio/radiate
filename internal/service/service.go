package service

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jaksonkallio/radiate/internal/service/graph"
	"github.com/jaksonkallio/radiate/internal/service/graph/generated"
)

type Service struct {
	gin *gin.Engine
}

func NewService() (*Service, error) {
	service := &Service{}
	err := service.Init()

	if err != nil {
		return nil, err
	}

	return service, nil
}

func (service *Service) Init() error {
	log.Println("Initializing service")

	service.gin = gin.Default()

	service.gin.POST("/query", graphqlHandler())
	service.gin.GET("/", playgroundHandler())

	return nil
}

func (service *Service) Serve() {
	http.ListenAndServe(":5011", service.gin)
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

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
