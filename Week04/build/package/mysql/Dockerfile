FROM mysql:8.0

ADD conf/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf

ADD moive.sql /docker-entrypoint-initdb.d/moive.sql

ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d

RUN mkdir -p /var/log/mysql-binlog && chown -R mysql:mysql /var/log/mysql-binlog
RUN mkdir -p /var/log/docker_log/mysql/ && chown -R mysql:mysql /var/log/docker_log/mysql/

ENV MYSQL_ROOT_PASSWORD=123456
ENV LANG C.UTF-8

EXPOSE 3306