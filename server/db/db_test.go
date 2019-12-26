package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "hostels",
		User:       "postgres",
		Password:   "1861",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://stanislav7766:1861@localhost/hostels?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
