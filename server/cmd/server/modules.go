//+build wireinject

package main

import (
	"github.com/stanislav7766/architecture-lab2/server/hostels"
	"github.com/google/wire"
)

// ComposeApiServer will create an instance of HostelApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*HostelApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go)
		NewDbConnection,
		// Add providers from channels package
		hostels.Providers,
		// Provide HostelApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(HostelApiServer), "Port", "HostelsHandler"),
	)
	return nil, nil
}
