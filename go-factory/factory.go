package factory

// # Abstract Factory Pattern #

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) string
}

type StripeProcessor struct{}

func (s *StripeProcessor) Process(amount float64) string {
	return fmt.Sprintf("Processing Stripe payment of $%.2f", amount)
}

type PaypalProcessor struct{}

func (p *PaypalProcessor) Process(amount float64) string {
	return fmt.Sprintf("Processing Paypal payment of $%.2f", amount)
}

type PaymentFactory interface {
	CreateProcessor() PaymentProcessor
}

type StripeFactory struct{}

func (s *StripeFactory) CreateProcessor() PaymentProcessor {
	return &StripeProcessor{}
}

type PaypalFactory struct{}

func (p *PaypalFactory) CreateProcessor() PaymentProcessor {
	return &PaypalProcessor{}
}

func NewPaymaneClient(factory PaymentFactory) PaymentProcessor {
	return factory.CreateProcessor()
}
