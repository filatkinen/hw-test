FROM postgres
ENV  POSTGRES_PASSWORD=pass
COPY  ./tests/integration-tests/pgsql/postgres_dump.sql /docker-entrypoint-initdb.d/
EXPOSE 5432
