package usecase

import "github.com/hugovallada/go-intensivo/internal/order/entity"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewGetTotalUseCase(orderRepository entity.OrderRepository) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()

	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
