FROM postgres:15

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=uof_db

# Copy SQL init file
COPY init.sql /docker-entrypoint-initdb.d/
