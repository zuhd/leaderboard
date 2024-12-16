package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type authRequest struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func (s *ServiceImpl) Auth(c *gin.Context) {
	var form authRequest
	httpCode, err := s.request(c, &form)
	if err != nil {
		s.response(c, httpCode, err, nil)
		return
	}

	ret, err := s.b.Auth(s.ctx, form.Username, form.Password)
	if err != nil {
		s.response(c, http.StatusInternalServerError, err, nil)
		return
	}

	s.response(c, http.StatusOK, nil, ret)
}
