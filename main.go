package main

import (
	"fmt"

	"cris92.com/go-patterns/go-decorator"
	"cris92.com/go-patterns/go-factory"
	"cris92.com/go-patterns/go-observer"
	"cris92.com/go-patterns/go-singleton"
)

func main() {
	fmt.Println("Hello, This is Go Patterns!")

	//  Singleton
	fmt.Println("Start testing Singleton Pattern...")
	singleton.TestConcurrency()
	fmt.Println("End testing Singleton Pattern...")

	fmt.Println("\n")

	//  Factory
	fmt.Println("Start testing Factory Pattern...")
	stripeFactory := factory.StripeFactory{}
	var stripeClient factory.PaymentProcessor

	stripeClient = factory.NewPaymaneClient(&stripeFactory)
	stripePaymentResult := stripeClient.Process(10.5)
	println(stripePaymentResult)
	fmt.Println("End testing Factory Pattern...")

	fmt.Println("\n")

	//  Observer
	fmt.Println("Start testing Observer Pattern...")
	stockTicker := observer.StockTicker{}

	investor1 := observer.Investor{
		Name: "Long",
	}

	investor2 := observer.Investor{
		Name: "Cris",
	}

	stockTicker.Attach(&investor1)
	stockTicker.Attach(&investor2)

	stockTicker.Notify("Hello, Observer! Your stock X has increased 1$!")
	fmt.Println("End testing Observer Pattern...")

	fmt.Println("\n")

	//  Decorator
	fmt.Println("Start testing Decorator Pattern...")
	composite := decorator.CompositeDecorator()
	authRes := composite("authenticated")
	unAuthRes := composite("anonymous")
	fmt.Println(authRes)
	fmt.Println(unAuthRes)
	fmt.Println("End testing Decorator Pattern...")

	fmt.Println("\n")

}
