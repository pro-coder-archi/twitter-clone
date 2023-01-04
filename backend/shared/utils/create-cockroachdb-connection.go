package utils

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateCockroachDBConnection(ctx context.Context) *sql.DB {

	uri := GetEnv("COCKROACHDB_URL")

	//* connecting to the cockroachDB instance
	dbConnection, error := sql.Open("postgres", uri)

	if(error != nil) {
		log.Fatalf("❌ error connecting to cockroachDB instance : ")

		log.Fatalf(error.Error( )) }

	defer dbConnection.Close( )

	error= dbConnection.Ping( )

	if(error != nil) {
		log.Fatalf("❌ error pinging cockroachDB instance : ")

		log.Fatalf(error.Error( )) }

	log.Println("🔥 successfully connected to cockroach database")

	return dbConnection
}