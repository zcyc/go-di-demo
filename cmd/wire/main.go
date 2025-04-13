package main

import (
	"fmt"
	"go-di-demo/pkg/wire"
)

func main() {
	fmt.Println("--- Wire 示例 ---")

	// 使用 Wire 生成的代码初始化依赖树
	userManager, err := wire.InitializeUserManager()
	if err != nil {
		fmt.Printf("初始化用户管理器错误: %v\n", err)
		return
	}

	// 使用包含所有注入依赖的 userManager
	userId := "user789"
	userData := "王五, wang@example.com"

	// 注册新用户
	err = userManager.RegisterUser(userId, userData)
	if err != nil {
		fmt.Printf("注册用户错误: %v\n", err)
		return
	}

	// 获取用户数据
	userDetails, err := userManager.GetUserDetails(userId)
	if err != nil {
		fmt.Printf("获取用户详情错误: %v\n", err)
		return
	}

	fmt.Printf("用户详情: %s\n", userDetails)
}
