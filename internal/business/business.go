package business

import (
	"fmt"
	"leaderboard/internal/dao"
)

type Business interface {
	rankBusiness
	authBusiness
	scoreBusiness
}

type genericBusiness struct {
	db dao.DAO
}

func NewBusiness(dsn string) (*genericBusiness, error) {
	var err error
	b := &genericBusiness{}
	b.db, err = dao.NewDAO(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to new dao: %w", err)
	}

	return b, nil
}
