FROM postgres:latest
ENV POSTGRES_USER=root \
    POSTGRES_PASSWORD=password \
    POSTGRES_DB=mydb
COPY init.sql /docker-entrypoint-initdb.d/
VOLUME /var/lib/postgresql/data
EXPOSE 5432
CMD ["postgres"]
