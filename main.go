package main

import (
	"leaderboard/internal/service"
)

func main() {
	s, err := service.NewService()
	if err != nil {
		panic("failed to new service")
	}

	err = s.Build()
	if err != nil {
		panic("failed to build service")
	}

	err = s.Serve()
	if err != nil {
		panic("failed to run service")
	}
}
