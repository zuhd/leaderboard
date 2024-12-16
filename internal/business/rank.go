package business

import (
	"context"
	"fmt"
	"leaderboard/internal/dao"
	error2 "leaderboard/internal/error"
	"time"
)

type PlayerRankModel struct {
	PlayerID   int64 `json:"player_id"`
	PlayerRank int64 `json:"player_rank"`
	UpdateTime int64 `json:"update_time"`
}

type rankBusiness interface {
	GetPlayerRankByID(ctx context.Context, playerID int64) (*PlayerRankModel, error)
	ListPlayerRank(ctx context.Context) ([]*PlayerRankModel, error)
	UpdatePlayerRankByID(ctx context.Context, playerID int64, playerRank int64) error
}

func playerRankDB2Model(db *dao.PlayerRank) *PlayerRankModel {
	return &PlayerRankModel{
		PlayerID:   db.PlayerID,
		PlayerRank: db.PlayerRank,
		UpdateTime: db.UpdateTime,
	}
}
func (b *genericBusiness) GetPlayerRankByID(ctx context.Context, playerID int64) (*PlayerRankModel, error) {
	if playerID <= 0 {
		return nil, error2.ErrInvalidPlayerID
	}

	return b.getPlayerRankByID(ctx, playerID)
}

func (b *genericBusiness) getPlayerRankByID(ctx context.Context, playerID int64) (*PlayerRankModel, error) {
	rank, err := b.db.GetPlayerRankByID(ctx, playerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get player rank: %w", err)
	}

	return playerRankDB2Model(rank), nil
}

func (b *genericBusiness) ListPlayerRank(ctx context.Context) ([]*PlayerRankModel, error) {
	ranks, err := b.db.ListPlayerRank(ctx)
	if err != nil {
		return []*PlayerRankModel{}, fmt.Errorf("failed to list player rank: %w", err)
	}

	models := make([]*PlayerRankModel, 0)
	for _, v := range ranks {
		models = append(models, playerRankDB2Model(v))
	}

	return models, nil
}

func (b *genericBusiness) UpdatePlayerRankByID(ctx context.Context, playerID int64, playerRank int64) error {
	if playerID <= 0 {
		return error2.ErrInvalidPlayerID
	}

	return b.updatePlayerRankByID(ctx, playerID, playerRank)
}

func (b *genericBusiness) updatePlayerRankByID(ctx context.Context, playerID int64, playerRank int64) error {
	now := time.Now().Unix()
	err := b.db.UpdatePlayerRankByID(ctx, playerID, playerRank, now)
	if err != nil {
		return fmt.Errorf("failed to update player rank: %w", err)
	}

	return nil
}
