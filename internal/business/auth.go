package business

import (
	"context"
	"fmt"
	error2 "leaderboard/internal/error"
	"leaderboard/internal/utils"
)

type AuthTokenModel struct {
	Token string `json:"token"`
}

type authBusiness interface {
	Auth(ctx context.Context, userName string, passWord string) (*AuthTokenModel, error)
}

func (b *genericBusiness) Auth(ctx context.Context, userName string, passWord string) (*AuthTokenModel, error) {
	if len(userName) == 0 || len(passWord) == 0 {
		return nil, error2.ErrInvalidUser
	}

	return b.auth(ctx, userName, passWord)
}

func (b *genericBusiness) auth(ctx context.Context, userName string, passWord string) (*AuthTokenModel, error) {
	err := b.db.Auth(ctx, userName, passWord)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth: %w", err)
	}

	token, err := utils.GenerateToken(userName, passWord)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &AuthTokenModel{Token: token}, nil
}
