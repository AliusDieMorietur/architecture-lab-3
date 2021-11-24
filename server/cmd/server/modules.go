//go:build wireinject
// +build wireinject

package main

import (
	"SamuraiLab3/server/data"

	"github.com/google/wire"
)

func ComposeApiServer(port int) (*ForumApiServer, error) {
	wire.Build(
		data.Providers,
		wire.Struct(new(ForumApiServer), "Port", "HandlerFunc"),
	)
	return nil, nil
}
