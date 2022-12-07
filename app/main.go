package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vandenbill/brand-commerce-product-query-service/repository/elasticsearch"
	"github.com/vandenbill/brand-commerce-product-query-service/usecase"
	"github.com/vandenbill/brand-commerce-product-query-service/util"
)

// TODO clean repo layer
// TODO implement logging
// TODO implement unit testing
// TODO sleep kidz!!!

func main() {
	app := fiber.New()

	closer := util.ConfigureJaeger()
	defer closer.Close()

	rabbitMQConn := util.DialRabbitMQ()
	defer rabbitMQConn.Close()

	ch, err := rabbitMQConn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q := util.DeclareQueue(ch)

	esClient, res := util.ElasticSearchClient()
	defer res.Body.Close()

	app.Get("/api/product", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	elasticSearchRepo := elasticsearch.NewElasticSearchRepo(esClient)
	productUsecase := usecase.NewProductUsecase(elasticSearchRepo)
	go func() {
		productUsecase.MonitorData(ch, q)
	}()

	log.Println("APP WORK")
	log.Fatal(app.Listen(":3000"))
}
