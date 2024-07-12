#!/bin/bash
source .env

export MIGRATION_DSN="host=pgsql port=$PG_PORT dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"
sleep 2 && goose -dir "./migrations" postgres "${MIGRATION_DSN}" up -v