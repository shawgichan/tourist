#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migrations -database "$DB_ADDRESS" -verbose up

echo "start the app"
exec "$@"