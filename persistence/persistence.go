package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DBSettings contains the settings of the DB
type DBSettings struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type Persistence struct {
	dbSettings *DBSettings
	database   *sql.DB
}

func NewPersistence(dbSettings *DBSettings) *Persistence {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbSettings.Host, dbSettings.Port, dbSettings.Username, dbSettings.Password, dbSettings.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}


	return &Persistence{
		dbSettings: dbSettings,
		database: db,
	}
}
