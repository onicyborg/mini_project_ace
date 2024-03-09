package database

import "database/sql"

var (
	DbConnection *sql.DB
)

func DbConnectionFunc(dbParam *sql.DB) {
	DbConnection = dbParam
}
