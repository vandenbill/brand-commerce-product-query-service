package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/vandenbill/brand-commerce-product-query-service/model/domain"
	"log"
	"time"
)

type elasticSearchRepo struct {
	esClient *elasticsearch.Client
}

func NewElasticSearchRepo(esClient *elasticsearch.Client) domain.ElasticSearchRepo {
	return &elasticSearchRepo{esClient: esClient}
}

func (e *elasticSearchRepo) SaveProduct(data map[string]interface{}, jaegerCtx context.Context) {
	id := fmt.Sprintf("%s", data["id"])
	jsonData, _ := json.Marshal(data)
	req := esapi.CreateRequest{
		Index:      "product-service",
		DocumentID: id,
		Body:       bytes.NewReader(jsonData),
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	res, err := req.Do(ctx, e.esClient)
	if err != nil {
		log.Panicf("IndexRequest ERROR: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("%s ERROR indexing document ID=%d", res.Status(), id)
	} else {

		var resMap map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
			log.Panicf("Error parsing the response body: %s", err)
		} else {
			fmt.Println("Status:", res.Status())
		}
	}
}

func (e *elasticSearchRepo) EditProduct(data map[string]interface{}, jaegerCtx context.Context) {
	id := fmt.Sprintf("%s", data["id"])
	jsonData, _ := json.Marshal(data)

	req := esapi.UpdateRequest{
		Index:      "product-service",
		DocumentID: id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, jsonData))),
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	res, err := req.Do(ctx, e.esClient)
	if err != nil {
		log.Panicf("IndexRequest ERROR: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("%s ERROR indexing document ID=%d", res.Status(), id)
	} else {

		var resMap map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
			log.Panicf("Error parsing the response body: %s", err)
		} else {
			fmt.Println("Status:", res.Status())
		}
	}
}

func (e *elasticSearchRepo) RemoveProduct(id interface{}, jaegerCtx context.Context) {
	stringID := fmt.Sprintf("%s", id)
	req := esapi.DeleteRequest{
		Index:      "product-service",
		DocumentID: stringID,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := req.Do(ctx, e.esClient)
	if err != nil {
		log.Panicf("IndexRequest ERROR: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("%s ERROR indexing document ID=%d", res.Status(), id)
	} else {

		var resMap map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
			log.Panicf("Error parsing the response body: %s", err)
		} else {
			fmt.Println("Status:", res.Status())
		}
	}
}
