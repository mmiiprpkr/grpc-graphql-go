package resolver

import "grpcgqlgo/generated/proto/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserServiceClient    services.UserServiceClient
	ProductServiceClient services.ProductServiceClient
}
