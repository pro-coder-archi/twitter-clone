package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateCockroachDBConnection( ) *sql.DB {
	const methodName= "CreateCockroachDBConnection"

	uri := GetEnv("COCKROACHDB_URL")

	//* connecting to the cockroachDB instance
	dbConnection, error := sql.Open("postgres", uri)

	if(error != nil) {
		Log(LogDetails{

			Method: methodName,
			Message: fmt.Sprintf("❌ error connecting to cockroachDB instance : \n%s", error.Error( )),
		})}

	error= dbConnection.Ping( )

	if(error != nil) {
		Log(LogDetails{

			Message: fmt.Sprintf("❌ error pinging cockroachDB instance : \n%s", error.Error( )),
			Method: methodName,
		})}

	Log(LogDetails{

		Type: DEFAULT_LOG,
		Method: methodName,
		Message: "🔥 successfully connected to cockroach database",
	})

	return dbConnection
}