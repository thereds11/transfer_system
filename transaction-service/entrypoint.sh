#!/bin/bash
set -e

echo "⏳ Waiting for PostgreSQL to be available..."
until pg_isready -h postgres -p 5432 -U transferuser > /dev/null 2>&1; do
  sleep 1
done

echo "✅ PostgreSQL is ready"

echo "⚙️ Running DB migrations..."
for file in db/migrations/*.sql; do
    echo "▶️ Running migration: $file"
    psql "$DB_URL" < "$file" || echo "⚠️ Migration $file failed or already applied, continuing..."
done

echo "🚀 Starting Transaction Service..."
exec ./main
