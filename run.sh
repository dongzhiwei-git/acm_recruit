docker build -t acm .
docker run -p 8000:8000 acm
docker run --name acm_mysql -v /opt/mysql/conf.d:/etc/mysql/conf.d -v /opt/mysql/logs:/logs -v /opt/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123 -di -p 8000:8000 mysql
