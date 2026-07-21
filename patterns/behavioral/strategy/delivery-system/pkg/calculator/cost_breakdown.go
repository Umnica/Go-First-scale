package calculator

// # Детализация стоимости

type CostBreakdown struct {
	BaseCost     float64
	DistanceCost float64
	WeightCost   float64
	ExtraCost    float64
	Total        float64
	Description  string
}
