package utils

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateDbConnection(ctx context.Context) {

	uri, isEnvFound := os.LookupEnv("DATABASE_URL")

	if(!isEnvFound) {
		log.Fatal("❌ env DATABASE_URL not found") }

	//* connecting to the database
	dbConnection, error := sql.Open("postgres", uri)

	if(error != nil) {
		log.Fatalf("❌ error connecting to cockroachDB instance : ")

		log.Fatalf(error.Error( )) }

	defer dbConnection.Close( )

	error= dbConnection.Ping( )

	if(error != nil) {
		log.Fatalf("❌ error pinging cockroachDB instance : ")

		log.Fatalf(error.Error( )) }

	log.Println("🔥 successfully connected to authentication database")
}