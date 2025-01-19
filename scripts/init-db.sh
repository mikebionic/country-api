#!/bin/bash
# Load environment variables from .env file if it exists
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

DB_HOST="${DB_HOST:-country_db}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-country_db}"
DB_PASSWORD="${DB_PASSWORD:-country_db}"
DB_NAME="${DB_NAME:-country_db}"
DB_SCHEMASDIR=$(pwd)/"${DB_SCHEMASDIR:schemas}"

echo "Checking DB connection.........."

PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -c "SELECT 1 FROM information_schema.tables WHERE table_name = 'tbl_country';" | grep -q 1

if [ $? -ne 0 ]; then
    echo "Starting DB initialization."

    PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -f $DB_SCHEMASDIR/0.0.2_create_country_city.sql
    echo "Initialization completed."
else
    echo "+++ DB already initialized. Skipping. +++"
fi
