#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	CREATE DATABASE customers_db;
	GRANT ALL PRIVILEGES ON DATABASE customers_db TO postgres;
EOSQL
