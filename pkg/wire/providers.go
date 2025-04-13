//go:build wireinject
// +build wireinject

package wire

import (
	"go-di-demo/pkg/common/dao"
	"go-di-demo/pkg/common/service"

	"github.com/google/wire"
)

var UserDAOSet = wire.NewSet(
	dao.NewInMemoryUserDAO,
	wire.Bind(new(dao.UserDAO), new(*dao.InMemoryUserDAO)),
)

var ProductDAOSet = wire.NewSet(
	dao.NewInMemoryProductDAO,
	wire.Bind(new(dao.ProductDAO), new(*dao.InMemoryProductDAO)),
)

// ProviderSet 是包含所有依赖的 Wire 提供者集合
var ProviderSet = wire.NewSet(
	UserDAOSet,
	ProductDAOSet,
	service.NewUserService,
	service.NewProductService,
	service.NewNotificationService,
	service.NewUserManager,
)

// InitializeUserManager 创建一个包含所有依赖的 UserManager
func InitializeUserManager() (*service.UserManager, error) {
	wire.Build(ProviderSet)
	return nil, nil // Wire 会忽略此返回；这只是为了使编译器满意
}
