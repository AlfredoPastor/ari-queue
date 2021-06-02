//+build wireinject

package internal

import "github.com/google/wire"

var superSet = wire.NewSet(
	NewServer,
)

func InitializeServer() (Server, error) {
	wire.Build(superSet)
	return Server{}, nil
}
