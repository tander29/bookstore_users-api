# bookstore_users-api

### Udemy Course following
golang-how-to-design-and-build-rest-microservices-in-go

### To Run App:

#### Env Variables:

Need to set env variables for Db connection: Check datasources/mysql for full list

mysql_users_username

mysql_users_password

mysql_users_host

mysql_users_schema


### Running MySql DB in Docker

1. `docker pull mysql`
2. create a mysql docker container and establish a host/port.  3306 is the default Mysql port within the docker container
i.e `docker run --env MYSQL_ROOT_PASSWORD=password -d -p 127.0.0.1:3307:3306 mysqlcontainername`
3. `sudo docker exec -it mysqlcontainername bash`
4. (should be in the docker container interactive terminal now, to work within the db:)`mysql -u root -p`

sudo docker exec -it mysqlcontainername bash