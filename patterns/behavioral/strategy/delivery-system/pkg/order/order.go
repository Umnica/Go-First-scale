package order

import (
	"fmt"
	"patterns/behavioral/strategy/delivery-system/pkg/calculator"
	"time"

	"github.com/google/uuid"
)

// # Модель заказа
type DeliveryStrategy interface {
	// CalculateCost возвращает общую стоимость доставки
	CalculateCost(order *Order) float64

	// CalculateTime возвращает время доставки в минутах
	CalculateTime(order *Order) int

	// IsAvailable проверяет, доступна ли стратегия для данного заказа
	IsAvailable(order *Order) bool

	// GetBreakdown возвращает детализацию стоимости (для чеков)
	GetBreakdown(order *Order) calculator.CostBreakdown

	// Name возвращает название стратегии
	Name() string

	// Priority возвращает приоритет выполнения (для экспресса)
	Priority(order *Order) int
}

type Order struct {
	ID         string
	CreatedAt  time.Time
	Restaurant string
	Items      []string
	Quantity   int     // общее количество позиций
	Weight     float64 // вес заказа в кг
	Distance   float64 // расстояние до ресторана в км
	Weather    string  // "sunny", "rainy", "snowy"
	Priority   int     // 1-5 (5 - высший)
	Strategy   DeliveryStrategy
}

func (order *Order) NewOrder(restaurant string) *Order {
	return &Order{
		Restaurant: restaurant,
		CreatedAt:  time.Now(),
		ID:         uuid.New().String(),
	}
}


// Display возвращает краткую информацию о заказе
func (order *Order) Display() string {
	if order.Strategy == nil {
		return fmt.Sprintf("Заказ #%s [стратегия не выбрана]", order.ID)
	}

	return fmt.Sprintf(
		"Заказ #%s | %s | %s | Стоимость: %.2f ₽ | Время: %d мин | %s",
		order.ID,
		order.Restaurant,
		order.Strategy.Name(),
		order.CalculateDeliveryCost(),
		order.CalculateDeliveryTime(),
		map[bool]string{true: "✅ Доступен", false: "❌ Недоступен"}[order.IsDeliveryAvailable()],
	)
}



func (order *Order) ChangeDeliveryStrategy(deliveryStrategy DeliveryStrategy) {
	order.Strategy = deliveryStrategy
}

func (order *Order) CalculateDeliveryCost() float64 {
	if order.Strategy == nil {
		return 0.0
	}

	return order.Strategy.CalculateCost(order)
}

func (order *Order) CalculateDeliveryTime() int {
	if order.Strategy == nil {
		return 0
	}

	return order.Strategy.CalculateTime(order)
}

func (order *Order) IsDeliveryAvailable() bool {
	if order.Strategy == nil {
		return false
	}

	return order.Strategy.IsAvailable(order)
}

func (order *Order) GetBreakdown() calculator.CostBreakdown {
	if order.Strategy == nil {
		return calculator.CostBreakdown{
			Description: "Стратегия доставки не выбрана",
		}
	}

	return order.Strategy.GetBreakdown(order)
}


func (order *Order) GetDeliveryInfo() string {
	if order.Strategy == nil {
		return ""
	}

	return order.Strategy.Name()
}
