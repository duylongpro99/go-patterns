package observer

// # Event driven Model #

import "fmt"

type Observer interface {
	Update(message string)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(msg string)
}

type StockTicker struct {
	observers []Observer
	symbol    string
	price     float64
}

func (s *StockTicker) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *StockTicker) Detach(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *StockTicker) Notify(msg string) {
	for _, observer := range s.observers {
		observer.Update(msg)
	}
}

type Investor struct {
	Name string
}

func (i *Investor) Update(message string) {
	fmt.Printf("%s received message: %s\n", i.Name, message)
}
