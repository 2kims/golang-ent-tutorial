FROM mysql:latest

RUN chown -R mysql:root /var/lib/mysql/

ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD
ARG DB_ROOT_PASSWORD

ENV DB_NAME=$DB_NAME
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_ROOT_PASSWORD=$DB_ROOT_PASSWORD

ADD data.sql /etc/mysql/data.sql

RUN sed -i 's/DB_NAME/'$DB_NAME'/g' /etc/mysql/data.sql
RUN cp /etc/mysql/data.sql /docker-entrypoint-initdb.d

EXPOSE 3306