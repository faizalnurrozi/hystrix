package main

import "github.com/afex/hystrix-go/hystrix"

func main() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	hystrix.Go("my_command", func() error {
		// talk to other services
		return nil
	}, nil)
}
