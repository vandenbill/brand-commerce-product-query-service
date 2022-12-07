package util

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func ElasticSearchClient() (*elasticsearch.Client, *esapi.Response) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_SEARCH_URI"),
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	return esClient, res
}
