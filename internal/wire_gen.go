// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package internal

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeServer() (Server, error) {
	server := NewServer()
	return server, nil
}

// wire.go:

var superSet = wire.NewSet(
	NewServer,
)
