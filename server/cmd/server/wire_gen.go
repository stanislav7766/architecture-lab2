//go:generate wire
//+build !wireinject
package main

import (
	"github.com/stanislav7766/architecture-lab2/server/hostels"
)

// Injectors from modules.go:
func ComposeApiServer(port HttpPortNumber) (*HostelApiServer, error) {

	db, err := NewDbConnection()

	if err != nil {
		return nil, err
	}
	store := hostels.NewStore(db)

	httpHandlerFunc := hostels.HttpHandler(store)
	hostelApiServer := &HostelApiServer{
		Port:           port,
		HostelsHandler: httpHandlerFunc,
	}
	return hostelApiServer, nil
}
