# gshop
Gshop is an API for shopping websites

# start

## Connect to mysql database
If mysql is not available, maybe you can use Docker to extract the mysql image
`docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=db_example -e MYSQL_DATABASE=gshop mysql`

## Connect to redis database
If redis is not available, maybe you can use Docker to extract the redis image
`docker run -d --name redis -p 6379:6379 redis --requirepass "rdb_exmaple"`