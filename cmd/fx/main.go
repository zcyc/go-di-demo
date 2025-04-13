package main

import (
	"fmt"
	"go-di-demo/pkg/fx"

	fxlib "go.uber.org/fx"
)

func main() {
	fmt.Println("--- Fx Example ---")

	app := fxlib.New(
		// Include our module with all dependencies
		fx.Module,

		// Register lifecycle hooks
		fxlib.Invoke(fx.RegisterLifecycle),
	)

	app.Run()
}
