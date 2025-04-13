package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Go Dependency Injection Examples")
	fmt.Println("===============================")
	fmt.Println("\nThis project demonstrates three popular dependency injection frameworks in Go:")
	fmt.Println("1. Dig (by Uber)")
	fmt.Println("2. Fx (by Uber)")
	fmt.Println("3. Wire (by Google)")
	fmt.Println("\nRunning examples...\n")

	// Run each example
	runExample("dig")
	fmt.Println("\n")

	runExample("fx")
	fmt.Println("\n")

	runExample("wire")
}

func runExample(name string) {
	fmt.Printf("Running %s example:\n", name)
	fmt.Println("----------------------------------------")

	// Execute the example command
	cmd := exec.Command("go", "run", fmt.Sprintf("./cmd/%s/main.go", name))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running example: %v\n", err)
	}

	fmt.Println("----------------------------------------")
}
