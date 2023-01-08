package leaderboard

import "leaderboard/internal/utils"

func main() {
	s, err := utils.NewService()
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
