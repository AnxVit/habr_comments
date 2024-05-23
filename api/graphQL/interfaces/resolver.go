package interfaces

import (
	"github.com/AnxVit/ozon_1/internal/transport/graphql/api"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	HabrService api.IHabrService
}
