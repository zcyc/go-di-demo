//go:build wireinject
// +build wireinject

package wire

import (
	"go-di-demo/pkg/common"
	"go-di-demo/pkg/common/dao"
	"go-di-demo/pkg/common/service"

	"github.com/google/wire"
)

// Bind concrete implementations to interfaces
var LoggerSet = wire.NewSet(
	common.NewSimpleLogger,
	wire.Bind(new(common.Logger), new(*common.SimpleLogger)),
)

var UserDAOSet = wire.NewSet(
	dao.NewInMemoryUserDAO,
	wire.Bind(new(dao.UserDAO), new(*dao.InMemoryUserDAO)),
)

var ProductDAOSet = wire.NewSet(
	dao.NewInMemoryProductDAO,
	wire.Bind(new(dao.ProductDAO), new(*dao.InMemoryProductDAO)),
)

// ProviderSet is a Wire provider set with all the dependencies
var ProviderSet = wire.NewSet(
	LoggerSet,
	UserDAOSet,
	ProductDAOSet,
	service.NewUserService,
	service.NewProductService,
	service.NewNotificationService,
	service.NewUserManager,
)

// InitializeUserManager creates a UserManager with all its dependencies
func InitializeUserManager() (*service.UserManager, error) {
	wire.Build(ProviderSet)
	return nil, nil // Wire will ignore this; it's just to make the compiler happy
}
