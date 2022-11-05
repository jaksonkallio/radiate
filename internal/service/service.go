package service

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jaksonkallio/radiate/internal/config"
	"github.com/jaksonkallio/radiate/internal/service/graph"
	"github.com/jaksonkallio/radiate/internal/service/graph/generated"
	"github.com/rs/zerolog/log"
	"net/http"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

type Service struct {
	gin       *gin.Engine
	shellIPFS *ipfsapi.Shell
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
	service.initHandlers()

	// Create IPFS shell.
	log.Info().Msg("initializing IPFS shell")
	service.shellIPFS = ipfsapi.NewShell(config.CurrentConfig.IPFSHost)

	// Report IPFS shell status.
	version, commitSHA, err := service.shellIPFS.Version()
	log.Info().
		Bool("connected", err != nil).
		Str("ipfs_version", version).
		Str("ipfs_commit", commitSHA).
		Msg("IPFS status")

	return nil
}

func (service *Service) initHandlers() {
	log.Info().Msg("initializing service handlers")

	service.gin = gin.Default()

	// All static files are accessible at `/static`
	service.gin.Static("/static", "./frontend/dist")

	// The index file is accessible at the root
	service.gin.StaticFile("/", "./frontend/dist/index.html")

	service.gin.POST("/query", graphqlHandler())
	service.gin.GET("/playground", playgroundHandler())
}

func (service *Service) ServeAPI() {
	log.Info().
		Str("api_host", config.CurrentConfig.APIHost).
		Msg("serving API")

	err := http.ListenAndServe(config.CurrentConfig.APIHost, service.gin)
	if err != nil {
		log.Error().Err(err).Msg("listen and serve failed")
	}
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
