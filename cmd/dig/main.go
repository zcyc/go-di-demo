package main

import (
	"fmt"
	"go-di-demo/pkg/common/service"
	"go-di-demo/pkg/dig"
)

func main() {
	container := dig.BuildContainer()

	// 使用构建好的依赖图执行代码
	err := container.Invoke(func(userManager *service.UserManager) {
		// 自动构建了所有依赖的 userManager
		userId := "user123"
		userData := "张三, zhang@example.com"

		fmt.Println("--- Dig 示例 ---")

		// 注册新用户
		err := userManager.RegisterUser(userId, userData)
		if err != nil {
			fmt.Printf("注册用户错误: %v\n", err)
			return
		}

		// 获取用户详情
		userDetails, err := userManager.GetUserDetails(userId)
		if err != nil {
			fmt.Printf("获取用户详情错误: %v\n", err)
			return
		}

		fmt.Printf("用户详情: %s\n", userDetails)
	})

	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
}
