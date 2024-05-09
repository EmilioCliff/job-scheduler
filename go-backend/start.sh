#!/bin/sh

set -e

# echo "running migrate down"
# /app/migrate -path /app/migration -database postgresql://root:secret@localhost:5432/job-scheduler?sslmode=disable -verbose drop -f

echo "running migrations"
/app/migrate -path /app/migrations -database postgresql://root:secret@postgres1:5432/job-scheduler?sslmode=disable -verbose up

exec "$@"