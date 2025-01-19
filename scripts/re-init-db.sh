#!/bin/bash
# Re-initializing DB
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

DB_HOST="${DB_HOST:-country_db}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-country_db}"
DB_PASSWORD="${DB_PASSWORD:-country_db}"
DB_NAME="${DB_NAME:-country_db}"
DB_SCHEMASDIR=$(pwd)/"${DB_SCHEMASDIR:schemas}"


echo "Removing DB."
PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -U "$DB_USER" -f $DB_SCHEMASDIR/0.0.1_drop_db.sql

echo "Initializing other scripts"
bash "$(pwd)/scripts/init-db.sh"

echo "DB RE-Initialization completed."
