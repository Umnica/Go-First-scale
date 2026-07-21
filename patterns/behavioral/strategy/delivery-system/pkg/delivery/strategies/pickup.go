package strategies

import (
	"patterns/behavioral/strategy/delivery-system/pkg/calculator"
	"patterns/behavioral/strategy/delivery-system/pkg/order"
)

// # Самовывоз

type Pickup struct{}

// Name возвращает название стратегии
func (p Pickup) Name() string {
	return "Самовывоз"
}

// CalculateCost возвращает общую стоимость доставки
func (p Pickup) CalculateCost(order *order.Order) float64 {
	return 0.0
}

// CalculateTime возвращает время доставки в минутах
func (p Pickup) CalculateTime(order *order.Order) int {
	return 0
}

// IsAvailable проверяет, доступна ли стратегия для данного заказа
func (p Pickup) IsAvailable(order *order.Order) bool {
	return true
}

// GetBreakdown возвращает детализацию стоимости (для чеков)
func (p Pickup) GetBreakdown(order *order.Order) calculator.CostBreakdown {
	// Сделать
	return calculator.CostBreakdown{
		Description: "Самовывоз (бесплатно)",
	}
}

// Priority возвращает приоритет выполнения (для экспресса)
func (p Pickup) Priority(order *order.Order) int {
	return 1
}
