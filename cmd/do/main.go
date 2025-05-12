package main

import (
	"fmt"
	"go-di-demo/pkg/common/service"
	doContainer "go-di-demo/pkg/do"

	"github.com/samber/do"
)

func main() {
	// 构建依赖注入容器
	injector := doContainer.BuildContainer()

	// 从容器中获取 UserManager
	userManager, err := do.Invoke[*service.UserManager](injector)
	if err != nil {
		fmt.Printf("获取 UserManager 失败: %v\n", err)
		return
	}

	// 使用 UserManager
	userId := "user123"
	userData := "李四, li@example.com"

	fmt.Println("--- Do 示例 ---")

	// 注册新用户
	err = userManager.RegisterUser(userId, userData)
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

	// 展示如何使用 do 的另一种方式：直接调用服务
	fmt.Println("\n--- 直接调用服务示例 ---")
	notificationService := do.MustInvoke[*service.NotificationService](injector)
	notificationService.Notify(userId, "这是一条来自 Do 的直接通知")
}
