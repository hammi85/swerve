FROM mariadb:latest

MAINTAINER Jan Michalowsky <sejamich@googlemail.com>

RUN mkdir -p /docker-entrypoint-initdb.d
COPY wordpress.sql /docker-entrypoint-initdb.d/wordpress.sql
RUN chmod -R 777 /docker-entrypoint-initdb.d
