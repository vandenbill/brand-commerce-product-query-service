package util

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

func ElasticSearchClient() (*elasticsearch.Client, *esapi.Response) {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	return esClient, res
}
