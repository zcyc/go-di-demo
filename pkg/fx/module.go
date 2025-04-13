package fx

import (
	"context"
	"fmt"
	"go-di-demo/pkg/common"
	"go-di-demo/pkg/common/dao"
	"go-di-demo/pkg/common/service"

	"go.uber.org/fx"
)

// 绑定函数，将具体类型转换为接口
func AsLogger(logger *common.SimpleLogger) common.Logger {
	return logger
}

func AsUserDAO(dao *dao.InMemoryUserDAO) dao.UserDAO {
	return dao
}

func AsProductDAO(dao *dao.InMemoryProductDAO) dao.ProductDAO {
	return dao
}

// Module 为 Fx 应用程序提供所有依赖项
var Module = fx.Options(
	fx.Provide(
		// 提供具体实现
		common.NewSimpleLogger,
		// 提供绑定函数
		AsLogger,

		// 提供 DAO 层
		dao.NewInMemoryUserDAO,
		AsUserDAO,
		dao.NewInMemoryProductDAO,
		AsProductDAO,

		// 提供服务层
		service.NewUserService,
		service.NewProductService,
		service.NewNotificationService,
		service.NewUserManager,
	),
)

// RegisterLifecycle 注册生命周期钩子
func RegisterLifecycle(lifecycle fx.Lifecycle, userManager *service.UserManager, logger *common.SimpleLogger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Log("应用程序启动中")

			// 示例数据
			userId := "user456"
			userData := "李四, li@example.com"

			// 应用启动时注册用户
			err := userManager.RegisterUser(userId, userData)
			if err != nil {
				return err
			}

			// 获取用户数据确认已保存
			userDetails, err := userManager.GetUserDetails(userId)
			if err != nil {
				return err
			}

			fmt.Printf("用户注册成功: %s\n", userDetails)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Log("应用程序关闭中")
			return nil
		},
	})
}
