package repository

import (
	"fmt"
	"goapp/internal/core/logic/common"
	"goapp/internal/repository/database"
	"sync/atomic"
)

var (
	repository atomic.Pointer[common.Repository]
)

type Config struct {
	Database database.Config
}

func new(config *Config) (*common.Repository, error) {
	if config == nil {
		return nil, fmt.Errorf("config is nil")
	}
	_, err := database.Open(config.Database)
	if err != nil {
		return nil, err
	}
	repo := &common.Repository{}
	return repo, nil
}

func Initialize(config *Config) error {
	repo, err := new(config)
	if err != nil {
		return err
	}
	repository.Store(repo)
	common.GetRepository = repository.Load
	return nil
}
