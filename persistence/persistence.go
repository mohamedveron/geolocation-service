package persistence

import (
	"database/sql"
)

type Persistence struct {
	sourcePath string
	database   *sql.DB
}

func NewPersistence(sourcePath string) *Persistence {


	return &Persistence{
		sourcePath: sourcePath,
	}
}
