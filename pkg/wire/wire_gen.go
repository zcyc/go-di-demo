// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"go-di-demo/pkg/common/dao"
	"go-di-demo/pkg/common/service"
)

// Injectors from providers.go:

// InitializeUserManager 创建一个包含所有依赖的 UserManager
func InitializeUserManager() (*service.UserManager, error) {
	inMemoryUserDAO := dao.NewInMemoryUserDAO()
	userService := service.NewUserService(inMemoryUserDAO)
	notificationService := service.NewNotificationService()
	userManager := service.NewUserManager(userService, notificationService)
	return userManager, nil
}

// providers.go:

var UserDAOSet = wire.NewSet(dao.NewInMemoryUserDAO, wire.Bind(new(dao.UserDAO), new(*dao.InMemoryUserDAO)))

var ProductDAOSet = wire.NewSet(dao.NewInMemoryProductDAO, wire.Bind(new(dao.ProductDAO), new(*dao.InMemoryProductDAO)))

// ProviderSet 是包含所有依赖的 Wire 提供者集合
var ProviderSet = wire.NewSet(
	UserDAOSet,
	ProductDAOSet, service.NewUserService, service.NewProductService, service.NewNotificationService, service.NewUserManager,
)
