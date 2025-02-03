package usecase

import (
	"fmt"

	"github.com/math-schenatto/go-clean-arch/internal/entity"
	"github.com/math-schenatto/go-clean-arch/pkg/events"
)

type OrderListOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
	Total  int              `json:"total"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {

	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute(page, limit int) (OrderListOutputDTO, error) {
	fmt.Println("Chegou  no use case")
	orders, err := l.OrderRepository.FindAll(page, limit)
	if err != nil {
		return OrderListOutputDTO{}, err
	}

	total, err := l.OrderRepository.GetTotal()
	if err != nil {
		return OrderListOutputDTO{}, err
	}
	fmt.Println("TOTALLLLLL ", total)
	var output []OrderOutputDTO
	for _, order := range orders {
		output = append(output, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return OrderListOutputDTO{
		Orders: output,
		Total:  total,
	}, nil
}
