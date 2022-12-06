package usecase

import (
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"github.com/vandenbill/brand-commerce-product-query-service/model/domain"
	"github.com/vandenbill/brand-commerce-product-query-service/util"
)

type productUsecase struct {
	elasticSearchRepo domain.ElasticSearchRepo
}

func NewProductUsecase(repo domain.ElasticSearchRepo) domain.ProductUsecase {
	return &productUsecase{elasticSearchRepo: repo}
}

func (p *productUsecase) MonitorData(ch *amqp091.Channel, q amqp091.Queue) {
	trace, ctx := util.JegerTrace("ProductQueryService MonitorData")
	defer trace.Finish()

	msgs := util.RegisterConsumer(ch, q)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var data map[string]interface{}
			err := json.Unmarshal(d.Body, &data)
			util.FailOnError(err, "Cant unmarshal data from message broker")
			switch data["method"] {
			case "create":
				p.CreateProductUsecase(data, ctx)
			case "update":
				p.UpdateProductUsecase(data, ctx)
			case "delete":
				p.DeleteProductUsecase(data, ctx)
			}
		}
	}()
	<-forever
}

func (p *productUsecase) CreateProductUsecase(data map[string]interface{}, jaegerCtx context.Context) {
	trace, ctx := util.JegerTrace("ProductQueryService CreateProductUsecase")
	defer trace.Finish()

	p.elasticSearchRepo.SaveProduct(data, ctx)
}

func (p *productUsecase) UpdateProductUsecase(data map[string]interface{}, jaegerCtx context.Context) {
	trace, ctx := util.JegerTrace("ProductQueryService UpdateProductUsecase")
	defer trace.Finish()

	p.elasticSearchRepo.EditProduct(data, ctx)
}

func (p *productUsecase) DeleteProductUsecase(data map[string]interface{}, jaegerCtx context.Context) {
	trace, ctx := util.JegerTrace("ProductQueryService DeleteProductUsecase")
	defer trace.Finish()

	p.elasticSearchRepo.RemoveProduct(data["id"], ctx)
}
