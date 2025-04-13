package main

import (
	"fmt"
	"go-di-demo/pkg/fx"

	fxlib "go.uber.org/fx"
)

func main() {
	fmt.Println("--- Fx 示例 ---")

	app := fxlib.New(
		// 包含所有依赖的模块
		fx.Module,

		// 注册生命周期钩子
		fxlib.Invoke(fx.RegisterLifecycle),
	)

	app.Run()
}
