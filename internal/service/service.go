package service

import (
	"context"
	"fmt"
	"leaderboard/internal/business"
	errors2 "leaderboard/internal/error"
	"leaderboard/internal/utils"
	"net/http"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Service interface {
	Build() error
	Serve() error
}

type ServiceImpl struct {
	Logger       *zap.Logger
	Viper        *viper.Viper
	Gin          *gin.Engine
	b            business.Business
	ctx          context.Context
	moduleConfig ModuleConfig
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewService() (*ServiceImpl, error) {
	s := &ServiceImpl{
		Logger: zap.L(),
		Viper:  viper.New(),
		Gin:    gin.New(),
		ctx:    context.Background(),
	}

	err := s.buildConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to build config %w", err)
	}

	b, err := business.NewBusiness(s.moduleConfig.Module.Leaderboard.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to create business %w", err)
	}

	s.b = b
	return s, nil
}

func (s *ServiceImpl) Build() error {
	// init router
	s.initRounter()
	return nil
}

// middleware logger
func (s *ServiceImpl) logger() gin.HandlerFunc {
	// init the logger or other prepare stuff here
	return func(c *gin.Context) {
		// these code only works when request comes
		// print log here when the request travels here
		c.Next()
	}
}

// middleware timeout
func (s *ServiceImpl) timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		ch := make(chan struct{}, 1)
		go func() {
			c.Next()
			ch <- struct{}{}
		}()

		select {
		case <-ch:
			return
		case <-ctx.Done():
			c.AbortWithError(http.StatusRequestTimeout, ctx.Err())
			return
		}
	}
}

// middleware jwt
func (s *ServiceImpl) jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var data interface{}
		var err error

		token := c.GetHeader("token")
		if token == "" {
			err = errors2.ErrInvalidHeader
		} else {
			_, err1 := utils.ParseToken(token)
			if err1 != nil {
				err = fmt.Errorf("failed to parse token %w", err1)
			}
		}

		if err != nil {
			s.response(c, http.StatusInternalServerError, err, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (s *ServiceImpl) Serve() error {
	server := &http.Server{
		Addr:         s.moduleConfig.Module.Leaderboard.Listen,
		Handler:      s.Gin,
		ReadTimeout:  time.Second * time.Duration(s.moduleConfig.Module.Leaderboard.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(s.moduleConfig.Module.Leaderboard.WriteTimeout),
	}

	//log.Printf("[info] start http server listening %s", endPoint)
	s.Logger.Info("start http server")

	server.ListenAndServe()
	return nil
}

func (s *ServiceImpl) buildConfig() error {
	s.Viper.AddConfigPath(".")
	s.Viper.AddConfigPath("./config")
	s.Viper.AddConfigPath("./configs/config")
	s.Viper.SetConfigName("config")
	err := s.Viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	s.Viper.SetConfigName("secret")
	s.Viper.MergeInConfig()
	if err := s.Viper.Unmarshal(&s.moduleConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return nil
}

func (s *ServiceImpl) request(c *gin.Context, form interface{}) (int, error) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, errors2.ErrInvalidParams
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to valid %w", err)
	}
	if !check {
		return http.StatusBadRequest, errors2.ErrInvalidParams
	}

	return http.StatusOK, nil
}

func (s *ServiceImpl) response(c *gin.Context, httpCode int, err error, data interface{}) {
	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = "ok"
	}

	c.JSON(httpCode, &Response{
		Msg:  msg,
		Data: data,
	})
}
