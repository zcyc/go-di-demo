package wire

import "go-di-demo/pkg/common"

// This file exists so that the wire package is not excluded by build constraints.
// The actual implementation will be generated by the Wire tool.

// UserManagerProvider is a placeholder for documentation
type UserManagerProvider interface {
	// InitializeUserManager creates a UserManager with all its dependencies
	InitializeUserManager() (*common.UserManager, error)
}
