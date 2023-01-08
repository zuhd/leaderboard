package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
)

type Service interface {
	Build() error
	Serve() error
}

type ServiceImpl struct {
	Logger *zap.Logger
	Viper  *viper.Viper
	Gin    *gin.Engine
}

type ModuleConfig struct {
	Name  string `yaml:"name" mapstructure:"name"`
	Owner string `yaml:"owner" mapstructure:"owner"`
}

func NewService() (*ServiceImpl, error) {
	s := &ServiceImpl{
		Logger: zap.L(),
		Viper:  viper.New(),
		Gin:    gin.New(),
	}

	return s, nil
}

func (s *ServiceImpl) Build() error {
	return nil
}

func (s *ServiceImpl) Serve() error {
	server := &http.Server{
		//Addr:           endPoint,
		Handler: s.Gin,
		//ReadTimeout:    readTimeout,
		//WriteTimeout:   writeTimeout,
		//MaxHeaderBytes: maxHeaderBytes,
	}

	//log.Printf("[info] start http server listening %s", endPoint)

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
	return nil
}
