package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Go 依赖注入示例")
	fmt.Println("===============================")
	fmt.Println("\n本项目演示了Go中四个流行的依赖注入框架：")
	fmt.Println("1. Dig（由Uber开发）")
	fmt.Println("2. Wire（由Google开发）")
	fmt.Println("3. Do（由Samber开发）")
	fmt.Println("4. Fx（由Uber开发）")
	fmt.Println("\n运行示例...\n")

	// 运行每个示例
	runExample("dig")
	fmt.Println("\n")

	runExample("wire")
	fmt.Println("\n")

	runExample("do")
	fmt.Println("\n")

	runExample("fx")
}

func runExample(name string) {
	fmt.Printf("运行 %s 示例：\n", name)
	fmt.Println("----------------------------------------")

	// 执行示例命令
	cmd := exec.Command("go", "run", fmt.Sprintf("./cmd/%s/main.go", name))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("运行示例错误: %v\n", err)
	}

	fmt.Println("----------------------------------------")
}
