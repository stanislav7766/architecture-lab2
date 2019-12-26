package hostels

import "database/sql"

type Store struct {
	Db *sql.DB
}

type Hostel struct {
	Speciality string `json:"speciality"`
}

type SendStudent struct {
	Name         string `json:"name"`
	Speciality      string `json:"speciality"`
	Hostel   string `json:"hostel"`
}

type Response struct {
	Hostel map[string] interface{} `json:"hostel"`
}
