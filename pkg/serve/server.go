package serve

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bonaysoft/engra/apis/graph"
	"github.com/bonaysoft/engra/pkg/dict"
	"github.com/spf13/cobra"
)

const defaultPort = "8081"

func Run(cmd *cobra.Command, args []string) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dicts, err := dict.NewDict()
	if err != nil {
		return err
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Dict: dicts}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}
