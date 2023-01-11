source ./credentials/cockroachdb-cluster.credentials.env

migration_operation=$1
migration_path=$2

migrate \
    -database "$CONNECTION_URL" \
    -path $migration_path $migration_operation