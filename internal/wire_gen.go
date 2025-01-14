// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/trinhdaiphuc/togo/configs"
	"github.com/trinhdaiphuc/togo/internal/api"
	"github.com/trinhdaiphuc/togo/internal/api/v1"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
	"github.com/trinhdaiphuc/togo/internal/repository"
	"github.com/trinhdaiphuc/togo/internal/service"
	"github.com/trinhdaiphuc/togo/internal/service/task"
	"github.com/trinhdaiphuc/togo/internal/service/user"
)

// Injectors from wire.go:

func InitializeServer() (*api.Server, func(), error) {
	config, err := configs.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := infrastructure.NewDB(config)
	if err != nil {
		return nil, nil, err
	}
	userRepository := repository.NewUserRepository(db)
	userService := userservice.NewUserService(userRepository, config)
	taskRepository := repository.NewTaskRepository(db)
	taskService := taskservice.NewTaskService(taskRepository)
	serviceService := service.NewService(userService, taskService)
	userHandler := v1.NewUserHandler(serviceService)
	taskHandler := v1.NewTaskHandler(serviceService)
	server := api.NewServer(config, userHandler, taskHandler)
	return server, func() {
		cleanup()
	}, nil
}

// wire.go:

var (
	infrastructureSet = wire.NewSet(infrastructure.NewDB)
	repositorySet     = wire.NewSet(repository.NewTaskRepository, repository.NewUserRepository)
	serviceSet        = wire.NewSet(service.NewService, userservice.NewUserService, taskservice.NewTaskService)
	handlerSet        = wire.NewSet(v1.NewTaskHandler, v1.NewUserHandler)
)
