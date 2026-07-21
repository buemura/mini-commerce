-- Executed once by the postgres container on first boot
-- (mounted into /docker-entrypoint-initdb.d).
-- Postgres has no "CREATE DATABASE IF NOT EXISTS", so generate the statement
-- conditionally and let psql's \gexec run it.

SELECT 'CREATE DATABASE customer_service'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'customer_service')\gexec

SELECT 'CREATE DATABASE order_service'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'order_service')\gexec

SELECT 'CREATE DATABASE payment_service'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'payment_service')\gexec

SELECT 'CREATE DATABASE product_service'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'product_service')\gexec
