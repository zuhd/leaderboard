package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getRankRequest struct {
	PlayerID int64 `json:"player_id" valid:"Required;Min(1)"`
}

func (s *ServiceImpl) GetRank(c *gin.Context) {
	var form getRankRequest
	httpCode, err := s.request(c, &form)
	if err != nil {
		s.response(c, httpCode, err, nil)
		return
	}

	rank, err := s.b.GetPlayerRankByID(s.ctx, form.PlayerID)
	if err != nil {
		s.response(c, http.StatusInternalServerError, err, nil)
		return
	}

	s.response(c, http.StatusOK, nil, rank)
}

type listLeaderboardRequest struct {
}

func (s *ServiceImpl) ListLeaderboard(c *gin.Context) {
	var form listLeaderboardRequest
	httpCode, err := s.request(c, &form)
	if err != nil {
		s.response(c, httpCode, err, nil)
		return
	}

	leaderboards, err := s.b.ListPlayerRank(s.ctx)
	if err != nil {
		s.response(c, http.StatusInternalServerError, err, nil)
		return
	}

	s.response(c, http.StatusOK, nil, leaderboards)
}
