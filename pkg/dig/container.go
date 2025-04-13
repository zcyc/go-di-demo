package dig

import (
	"go-di-demo/pkg/common"

	"go.uber.org/dig"
)

// BuildContainer creates and configures a dig container with all dependencies
func BuildContainer() *dig.Container {
	container := dig.New()

	// Provide concrete types and bind them to interfaces
	container.Provide(common.NewSimpleLogger, dig.As(new(common.Logger)))
	container.Provide(common.NewInMemoryDB, dig.As(new(common.Database)))

	// Provide other services
	container.Provide(common.NewUserService)
	container.Provide(common.NewNotificationService)
	container.Provide(common.NewUserManager)

	return container
}
