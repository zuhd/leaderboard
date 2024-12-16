package service

import (
	"github.com/gin-gonic/gin"
	"time"
)

func (s *ServiceImpl) initRounter() {
	// global middleware chain
	s.Gin.Use(s.logger())
	s.Gin.Use(s.timeout(time.Second * 5))
	s.Gin.Use(gin.Recovery())

	// global routers
	s.Gin.POST("/auth", s.Auth)

	// be careful, a new group is created
	apiv1 := s.Gin.Group("/api/v1")
	//apiv1.Use(s.jwt())
	apiv1.POST("/getscore", s.GetScore)
	apiv1.POST("/addscore", s.AddScore)
	apiv1.POST("/resetscore", s.ResetScore)
	apiv1.POST("/getrank", s.GetRank)
	apiv1.POST("/list", s.ListLeaderboard)
}
