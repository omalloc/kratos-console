package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewConsoleService,
	NewZoneService,
	NewNodeService,
	NewAppService,
	NewDiscoveryService,
	NewNamespaceService,
	// User, Role, Permissions,
	NewUserService,
	NewRoleService,
	NewPermissionService,
)
