package dig

import (
	"go-di-demo/pkg/common/dao"
	"go-di-demo/pkg/common/service"

	"go.uber.org/dig"
)

// BuildContainer 创建并配置一个包含所有依赖的 dig 容器
func BuildContainer() *dig.Container {
	container := dig.New()

	// 提供 DAO 层
	container.Provide(dao.NewInMemoryUserDAO, dig.As(new(dao.UserDAO)))
	container.Provide(dao.NewInMemoryProductDAO, dig.As(new(dao.ProductDAO)))

	// 提供 Service 层
	container.Provide(service.NewUserService)
	container.Provide(service.NewProductService)
	container.Provide(service.NewNotificationService)
	container.Provide(service.NewUserManager)

	return container
}
