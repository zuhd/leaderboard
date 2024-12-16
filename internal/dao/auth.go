package dao

import (
	"context"
	"fmt"
	error2 "leaderboard/internal/error"
)

type authDAO interface {
	Auth(ctx context.Context, userName string, passWord string) error
}

func (d *genericDAO) Auth(ctx context.Context, userName string, passWord string) error {
	query, err := getQueryStatement(auth)
	if err != nil {
		return fmt.Errorf("failed to get query statement: %w", err)
	}

	var usernames []string
	err = d.db.SelectContext(ctx, &usernames, query, userName, passWord)
	if err != nil {
		return fmt.Errorf("failed to select ranks: %w", err)
	}

	if len(usernames) == 0 {
		return error2.ErrInvalidUser
	}

	return nil
}
