// Service logic.
//
// Credit to https://gqlgen.com/recipes/gin/ for the handler functions.
package service

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jaksonkallio/radiate/internal/service/graph"
	"github.com/jaksonkallio/radiate/internal/service/graph/generated"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

var ServerSideRenderedBundle = ""

type Service struct {
	gin       *gin.Engine
	shellIPFS *ipfsapi.Shell
}

func NewService(clientIPFS *ipfsapi.Shell) (*Service, error) {
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

	// All static files are accessible at `/static`
	service.gin.Static("/static", "./frontend/dist")

	// The index file is accessible at the root
	service.gin.StaticFile("/", "./frontend/dist/index.html")

	service.gin.POST("/query", graphqlHandler())
	service.gin.GET("/playground", playgroundHandler())

	return nil
}

func (service *Service) Serve() {
	http.ListenAndServe(":5011", service.gin)
}

// Defining the Graphql handler.
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler.
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
