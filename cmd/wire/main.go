package main

import (
	"fmt"
	"go-di-demo/pkg/wire"
)

func main() {
	fmt.Println("--- Wire Example ---")

	// Initialize the dependency tree using Wire-generated code
	userManager, err := wire.InitializeUserManager()
	if err != nil {
		fmt.Printf("Error initializing user manager: %v\n", err)
		return
	}

	// Use the userManager with all its injected dependencies
	userId := "user789"
	userData := "Bob Johnson, bob@example.com"

	// Register a new user
	err = userManager.RegisterUser(userId, userData)
	if err != nil {
		fmt.Printf("Error registering user: %v\n", err)
		return
	}

	// Retrieve user data
	userDetails, err := userManager.GetUserDetails(userId)
	if err != nil {
		fmt.Printf("Error getting user details: %v\n", err)
		return
	}

	fmt.Printf("User details: %s\n", userDetails)
}
