package decorator

import "fmt"

// # HTTP Middleware Simulation #

type Handler func(request string) string

func LogDecorator(next Handler) Handler {
	return func(request string) string {
		fmt.Printf("Receive request: %s\n", request)
		return next(request)
	}

}

func AuthDecorator(next Handler) Handler {
	return func(request string) string {
		if request != "authenticated" {
			return "Unauthorized"
		}

		return next(request)
	}
}

func BasicHandler(request string) string {
	return fmt.Sprintf("Processed %s", request)

}

func CompositeDecorator() Handler {
	return LogDecorator(AuthDecorator(BasicHandler))
}
