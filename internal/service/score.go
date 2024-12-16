package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getScoreRequest struct {
	PlayerID int64 `json:"player_id" valid:"Required; Min(1)"`
}

func (s *ServiceImpl) GetScore(c *gin.Context) {
	var form getScoreRequest
	httpCode, err := s.request(c, &form)
	if err != nil {
		s.response(c, httpCode, err, nil)
		return
	}

	score, err := s.b.GetPlayerScoreByID(s.ctx, form.PlayerID)
	if err != nil {
		s.response(c, http.StatusInternalServerError, err, nil)
		return
	}

	s.response(c, http.StatusOK, nil, score)
}

type addScoreRequest struct {
	PlayerID int64 `json:"player_id" valid:"Required; Min(1)"`
	Score    int64 `json:"score" valid:"Required; Min(1)"`
}

func (s *ServiceImpl) AddScore(c *gin.Context) {
	var form addScoreRequest
	httpCode, err := s.request(c, &form)
	if err != nil {
		s.response(c, httpCode, err, nil)
		return
	}

	err = s.b.AddPlayerScoreByID(s.ctx, form.PlayerID, form.Score)
	if err != nil {
		s.response(c, http.StatusInternalServerError, err, "failed")
		return
	}

	s.response(c, http.StatusOK, nil, "success")
}

type resetScoreRequest struct {
	PlayerID int64 `json:"player_id" valid:"Required; Min(1)"`
}

func (s *ServiceImpl) ResetScore(c *gin.Context) {
	var form resetScoreRequest
	httpCode, err := s.request(c, &form)
	if err != nil {
		s.response(c, httpCode, err, nil)
		return
	}

	err = s.b.ResetPlayerScoreByID(s.ctx, form.PlayerID)
	if err != nil {
		s.response(c, http.StatusInternalServerError, err, "failed")
		return
	}

	s.response(c, http.StatusOK, nil, "success")
}
