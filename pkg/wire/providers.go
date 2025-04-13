//go:build wireinject
// +build wireinject

package wire

import (
	"go-di-demo/pkg/common"

	"github.com/google/wire"
)

// Bind concrete implementations to interfaces
var LoggerSet = wire.NewSet(
	common.NewSimpleLogger,
	wire.Bind(new(common.Logger), new(*common.SimpleLogger)),
)

var DatabaseSet = wire.NewSet(
	common.NewInMemoryDB,
	wire.Bind(new(common.Database), new(*common.InMemoryDB)),
)

// ProviderSet is a Wire provider set with all the dependencies
var ProviderSet = wire.NewSet(
	LoggerSet,
	DatabaseSet,
	common.NewUserService,
	common.NewNotificationService,
	common.NewUserManager,
)

// InitializeUserManager creates a UserManager with all its dependencies
func InitializeUserManager() (*common.UserManager, error) {
	wire.Build(ProviderSet)
	return nil, nil // Wire will ignore this; it's just to make the compiler happy
}
