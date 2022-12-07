package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vandenbill/brand-commerce-product-query-service/model/domain"
)

type SearchProductDelivery struct {
	productUsecase domain.ProductUsecase
}

func (s *SearchProductDelivery) SearchProductDelivery(c *fiber.Ctx) error {
	// TODO tuning search, so many configuration for search in elastic search

	return nil
}
