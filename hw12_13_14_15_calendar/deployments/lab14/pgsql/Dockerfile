FROM postgres
ENV  POSTGRES_PASSWORD=pass
COPY  ./deployments/lab14/pgsql/postgres_dump.sql /docker-entrypoint-initdb.d/
EXPOSE 5432
