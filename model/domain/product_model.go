package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
)

type ElasticSearchRepo interface {
	SaveProduct(data map[string]interface{}, jaegerCtx context.Context)
	EditProduct(data map[string]interface{}, jaegerCtx context.Context)
	RemoveProduct(id interface{}, jaegerCtx context.Context)
}

type ProductUsecase interface {
	MonitorData(ch *amqp091.Channel, q amqp091.Queue)
	CreateProductUsecase(data map[string]interface{}, jaegerCtx context.Context)
	UpdateProductUsecase(data map[string]interface{}, jaegerCtx context.Context)
	DeleteProductUsecase(data map[string]interface{}, jaegerCtx context.Context)
}

type ProductDelivery interface {
	SearchProductDelivery(c *fiber.Ctx) error
}
