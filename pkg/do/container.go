package do

import (
	"go-di-demo/pkg/common/dao"
	"go-di-demo/pkg/common/service"

	"github.com/samber/do"
)

// BuildContainer 构建依赖注入容器
func BuildContainer() *do.Injector {
	// 创建依赖注入容器
	injector := do.New()

	// 注册 UserDAO
	do.Provide(injector, func(i *do.Injector) (dao.UserDAO, error) {
		return dao.NewInMemoryUserDAO(), nil
	})

	// 注册 NotificationService
	do.Provide(injector, func(i *do.Injector) (*service.NotificationService, error) {
		return service.NewNotificationService(), nil
	})

	// 注册 UserService (依赖 UserDAO)
	do.Provide(injector, func(i *do.Injector) (*service.UserService, error) {
		userDAO := do.MustInvoke[dao.UserDAO](i)
		return service.NewUserService(userDAO), nil
	})

	// 注册 UserManager (依赖 UserService 和 NotificationService)
	do.Provide(injector, func(i *do.Injector) (*service.UserManager, error) {
		userService := do.MustInvoke[*service.UserService](i)
		notificationService := do.MustInvoke[*service.NotificationService](i)
		return service.NewUserManager(userService, notificationService), nil
	})

	return injector
}
