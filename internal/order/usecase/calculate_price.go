package usecase

import (
	"github.com/hugovallada/go-intensivo/internal/order/entity"
)

type OrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(orderInput OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(orderInput.ID, orderInput.Price, orderInput.Tax)

	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil

}
