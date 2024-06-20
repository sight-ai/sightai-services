package lib_subgraph

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/hasura/go-graphql-client"
)

var Clients []*Client

type Client struct {
	Node    *Node
	Graphql *graphql.Client
}

// Initialize the kafka connections based configuration
func Initialize(config *Config) {
	if config == nil {
		log.Fatal().Msg("subgraph config file is empty")
	}

	var err error
	for _, node := range config.Nodes {
		c := graphql.NewClient(node.Endpoint, nil)
		if err != nil {
			log.Fatal().Msg("failed top connect to subgraph")
		}
		Clients = append(Clients, &Client{
			Node:    node,
			Graphql: c,
		})
	}
}
