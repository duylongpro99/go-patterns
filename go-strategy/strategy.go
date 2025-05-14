package strategy

// # Ecommerce Promotion Calculation #

type PromotionStrategy interface {
	Calculate(amount float64) float64
}

// PercentDiscount Percentage discount strategy
type PercentageDiscount struct {
	Rate float64
}

func (p *PercentageDiscount) Calculate(amount float64) float64 {
	return amount * (1 - p.Rate)
}

// FixedDiscount fixed amount discount strategy
type FixedDiscount struct {
	Offset float64 //fixed reduction amount
}

func (p *FixedDiscount) Calculate(amount float64) float64 {
	if amount > p.Offset {
		return amount - p.Offset
	}

	return amount
}

// CheckoutContext checkout context
type CheckoutContext struct {
	Strategy PromotionStrategy
}

func (c *CheckoutContext) SetStrategy(strategy PromotionStrategy) {
	c.Strategy = strategy
}

func (c *CheckoutContext) CalculateFinalAmount(amount float64) float64 {
	return c.Strategy.Calculate(amount)
}
