package main

import (
	"fmt"
	"go-di-demo/pkg/common"
	"go-di-demo/pkg/dig"
)

func main() {
	container := dig.BuildContainer()

	// Execute code with the fully constructed dependency graph
	err := container.Invoke(func(userManager *common.UserManager) {
		// Use the userManager that was automatically constructed with all its dependencies
		userId := "user123"
		userData := "John Doe, john@example.com"

		fmt.Println("--- Dig Example ---")

		// Register a new user
		err := userManager.RegisterUser(userId, userData)
		if err != nil {
			fmt.Printf("Error registering user: %v\n", err)
			return
		}

		// Get user details
		userDetails, err := userManager.GetUserDetails(userId)
		if err != nil {
			fmt.Printf("Error getting user details: %v\n", err)
			return
		}

		fmt.Printf("User details: %s\n", userDetails)
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
