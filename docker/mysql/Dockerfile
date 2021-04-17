FROM mysql:5.7

ENV TZ Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && chown -R mysql:root /var/lib/mysql/

COPY ./docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./docker/mysql/docker-entrypoint-initdb.d/init.sql /docker-entrypoint-initdb.d/

CMD ["mysqld"]

EXPOSE 3306