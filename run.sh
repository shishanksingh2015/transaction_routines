#!/bin/bash

chmod +x run.sh

set -e

echo "Starting routines docker ..."
make run-docker

echo "Waiting for Database..."
until docker exec -it go_db pg_isready -U user -d transaction-routines > /dev/null 2>&1; do
  echo "Waiting for database to be ready..."
  sleep 2
done

echo "Database is ready!"

echo "Running migrations..."

make migrate-up

echo "Routines started successfully!"