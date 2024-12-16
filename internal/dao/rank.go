package dao

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PlayerRank struct {
	PlayerID   int64 `db:"player_id"`
	PlayerRank int64 `db:"player_rank"`
	UpdateTime int64 `db:"update_time"`
}

type rankDAO interface {
	GetPlayerRankByID(ctx context.Context, playerID int64) (*PlayerRank, error)
	GetPlayerRankByIDS(ctx context.Context, playerIDS []int64) ([]*PlayerRank, error)
	ListPlayerRank(ctx context.Context) ([]*PlayerRank, error)
	UpdatePlayerRankByID(ctx context.Context, playerID int64, rank int64, updateTime int64) error
}

func (d *genericDAO) GetPlayerRankByID(ctx context.Context, playerID int64) (*PlayerRank, error) {
	query, err := getQueryStatement(getPlayerRank)
	if err != nil {
		return nil, fmt.Errorf("failed to get query statement: %w", err)
	}

	var rank PlayerRank
	err = d.db.GetContext(ctx, &rank, query, playerID)
	if err != nil {
		return nil, fmt.Errorf("failed to select rank: %w", err)
	}

	return &rank, nil
}

func (d *genericDAO) GetPlayerRankByIDS(ctx context.Context, playerIDS []int64) ([]*PlayerRank, error) {
	query, err := getQueryStatement(getPlayersRanks)
	if err != nil {
		return []*PlayerRank{}, fmt.Errorf("failed to get query statement: %w", err)
	}

	var ranks []*PlayerRank
	query, args, err := sqlx.In(query, playerIDS)
	if err != nil {
		return []*PlayerRank{}, fmt.Errorf("failed to sqlx in: %w", err)
	}

	query = d.db.Rebind(query)
	err = d.db.SelectContext(ctx, &ranks, query, args...)
	if err != nil {
		return []*PlayerRank{}, fmt.Errorf("failed to select ranks: %w", err)
	}

	return ranks, nil
}

func (d *genericDAO) ListPlayerRank(ctx context.Context) ([]*PlayerRank, error) {
	query, err := getQueryStatement(listPlayerRank)
	if err != nil {
		return []*PlayerRank{}, fmt.Errorf("failed to get query statement: %w", err)
	}

	var ranks []*PlayerRank
	err = d.db.SelectContext(ctx, &ranks, query)
	if err != nil {
		return []*PlayerRank{}, fmt.Errorf("failed to select ranks: %w", err)
	}

	return ranks, nil
}

func (d *genericDAO) UpdatePlayerRankByID(ctx context.Context, playerID int64, rank int64, updateTime int64) error {
	query, err := getQueryStatement(updatePlayerRank)
	if err != nil {
		return fmt.Errorf("failed to get query statement: %w", err)
	}

	result, err := d.db.ExecContext(ctx, query, rank, updateTime, playerID)
	if err != nil {
		return fmt.Errorf("failed to exec context: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return errors.New("failed to update due to rows is zero")
	}

	return nil
}
