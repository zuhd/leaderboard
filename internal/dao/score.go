package dao

import (
	"context"
	"errors"
	"fmt"
)

type PlayerScore struct {
	PlayerID    int64 `db:"player_id"`
	PlayerScore int64 `db:"player_score"`
	UpdateTime  int64 `db:"update_time"`
}

type scoreDAO interface {
	GetPlayerScoreByID(ctx context.Context, playerID int64) (*PlayerScore, error)
	AddPlayerScoreByID(ctx context.Context, playerID int64, score int64, updateTime int64) error
	ResetPlayerScoreByID(ctx context.Context, playerID int64, updateTime int64) error
	ResetAllPlayerScore(ctx context.Context, updateTime int64) error
}

func (d *genericDAO) GetPlayerScoreByID(ctx context.Context, playerID int64) (*PlayerScore, error) {
	query, err := getQueryStatement(getPlayerScore)
	if err != nil {
		return nil, fmt.Errorf("failed to get query statement: %w", err)
	}

	var score PlayerScore
	err = d.db.GetContext(ctx, &score, query, playerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get score: %w", err)
	}

	return &score, nil
}

func (d *genericDAO) AddPlayerScoreByID(ctx context.Context, playerID int64, score int64, updateTime int64) error {
	query, err := getQueryStatement(addPlayerScore)
	if err != nil {
		return fmt.Errorf("failed to get query statement: %w", err)
	}

	result, err := d.db.ExecContext(ctx, query, score, updateTime, playerID)
	if err != nil {
		return fmt.Errorf("failed to add score: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (d *genericDAO) ResetPlayerScoreByID(ctx context.Context, playerID int64, updateTime int64) error {
	query, err := getQueryStatement(resetPlayerScore)
	if err != nil {
		return fmt.Errorf("failed to get query statement: %w", err)
	}

	result, err := d.db.ExecContext(ctx, query, updateTime, playerID)
	if err != nil {
		return fmt.Errorf("failed to reset score: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (d *genericDAO) ResetAllPlayerScore(ctx context.Context, updateTime int64) error {
	query, err := getQueryStatement(resetAllPlayerScore)
	if err != nil {
		return fmt.Errorf("failed to get query statement: %w", err)
	}

	result, err := d.db.ExecContext(ctx, query, updateTime)
	if err != nil {
		return fmt.Errorf("failed to reset all score: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
