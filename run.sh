#!/bin/bash

set -e

echo "Starting Docker containers..."
make run-docker

echo "Waiting for PostgreSQL to be ready..."
until docker exec -it go_db pg_isready -U user -d transaction-routines > /dev/null 2>&1; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 2
done

echo "PostgreSQL is ready!"

echo "Running migrations..."

make migrate-up

echo "Starting application..."

echo "Application started successfully!"