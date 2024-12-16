package business

import (
	"context"
	"fmt"
	"leaderboard/internal/dao"
	error2 "leaderboard/internal/error"
	"time"
)

type PlayerScoreModel struct {
	PlayerID    int64 `json:"player_id"`
	PlayerScore int64 `json:"score"`
	UpdateTime  int64 `json:"update_time"`
}

type scoreBusiness interface {
	GetPlayerScoreByID(ctx context.Context, playerID int64) (*PlayerScoreModel, error)
	AddPlayerScoreByID(ctx context.Context, playerID int64, score int64) error
	ResetPlayerScoreByID(ctx context.Context, playerID int64) error
}

func playerScoreDB2Model(db *dao.PlayerScore) *PlayerScoreModel {
	return &PlayerScoreModel{
		PlayerID:    db.PlayerID,
		PlayerScore: db.PlayerScore,
		UpdateTime:  db.UpdateTime,
	}
}

func (b *genericBusiness) GetPlayerScoreByID(ctx context.Context, playerID int64) (*PlayerScoreModel, error) {
	if playerID <= 0 {
		return nil, error2.ErrInvalidPlayerID
	}

	return b.getPlayerScoreByID(ctx, playerID)
}

func (b *genericBusiness) getPlayerScoreByID(ctx context.Context, playerID int64) (*PlayerScoreModel, error) {
	score, err := b.db.GetPlayerScoreByID(ctx, playerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get player score: %w", err)
	}

	return playerScoreDB2Model(score), nil
}

func (b *genericBusiness) AddPlayerScoreByID(ctx context.Context, playerID int64, score int64) error {
	if playerID <= 0 {
		return error2.ErrInvalidPlayerID
	}

	currentScore, err := b.getPlayerScoreByID(ctx, playerID)
	if err != nil {
		return fmt.Errorf("failed to get player score: %w", err)
	}

	afterScore := currentScore.PlayerScore + score
	if afterScore < 0 {
		return error2.ErrInvalidScore
	}
	return b.addPlayerScoreByID(ctx, playerID, afterScore)
}

func (b *genericBusiness) addPlayerScoreByID(ctx context.Context, playerID int64, score int64) error {
	now := time.Now().Unix()
	err := b.db.AddPlayerScoreByID(ctx, playerID, score, now)
	if err != nil {
		return fmt.Errorf("failed to add player score: %w", err)
	}

	return nil
}

func (b *genericBusiness) ResetPlayerScoreByID(ctx context.Context, playerID int64) error {
	if playerID <= 0 {
		return b.resetAllPlayerScore(ctx)
	}

	return b.resetAllPlayerScore(ctx)
}

func (b *genericBusiness) resetPlayerScoreByID(ctx context.Context, playerID int64) error {
	now := time.Now().Unix()
	err := b.db.ResetPlayerScoreByID(ctx, playerID, now)
	if err != nil {
		return fmt.Errorf("failed to reset player score: %w", err)
	}

	return nil
}

func (b *genericBusiness) resetAllPlayerScore(ctx context.Context) error {
	now := time.Now().Unix()
	err := b.db.ResetAllPlayerScore(ctx, now)
	if err != nil {
		return fmt.Errorf("failed to reset all player score: %w", err)
	}

	return nil
}
