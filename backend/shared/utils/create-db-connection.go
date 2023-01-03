package utils

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateDbConnection(ctx context.Context) {

	//* connecting to the database
	dbConnection, error := sql.Open("postgres", "postgresql://cluster-admin:EeS13BjMaDKHidpIpVxDuw@twitter-clone-cockroach-cluster-1806.7s5.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full")

	if(error != nil) {
		log.Fatalf("❌ error connecting to cockroachDB instance : ")

		log.Fatalf(error.Error( )) }

	defer dbConnection.Close( )

	error= dbConnection.Ping( )

	if(error != nil) {
		log.Fatalf("❌ error pinging cockroachDB instance : ")

		log.Fatalf(error.Error( )) }
}