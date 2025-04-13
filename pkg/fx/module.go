package fx

import (
	"context"
	"fmt"
	"go-di-demo/pkg/common"

	"go.uber.org/fx"
)

// Binding functions to convert concrete types to interfaces
func AsLogger(logger *common.SimpleLogger) common.Logger {
	return logger
}

func AsDatabase(db *common.InMemoryDB) common.Database {
	return db
}

// Module provides all dependencies for the Fx application
var Module = fx.Options(
	fx.Provide(
		// Provide concrete implementations
		common.NewSimpleLogger,
		// Provide binding functions
		AsLogger,
		// Provide DB that depends on Logger interface
		common.NewInMemoryDB,
		// Provide binding function for Database
		AsDatabase,
		// Provide other services
		common.NewUserService,
		common.NewNotificationService,
		common.NewUserManager,
	),
)

// RegisterLifecycle registers lifecycle hooks
func RegisterLifecycle(lifecycle fx.Lifecycle, userManager *common.UserManager, logger *common.SimpleLogger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Log("Starting application")

			// Example data
			userId := "user456"
			userData := "Jane Smith, jane@example.com"

			// Register a user on application start
			err := userManager.RegisterUser(userId, userData)
			if err != nil {
				return err
			}

			// Retrieve user data to confirm it was saved
			userDetails, err := userManager.GetUserDetails(userId)
			if err != nil {
				return err
			}

			fmt.Printf("User registered successfully: %s\n", userDetails)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Log("Stopping application")
			return nil
		},
	})
}
