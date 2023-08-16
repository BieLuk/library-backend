### Simple Library Backend
This repo contains Golang backend REST API using Gin framework. 
This simple application uses Viper library for handling configuration, Gorm as sql ORM for postgresql and goose as migration tool.
</br>
Application is containerised with docker-compose.

To run application:
1. run command `docker-compose up`.
2. After running application you need to migrate Postgresql DB schema. To do that you can run script `migrate_db.sh` from `tools` catalog (the script uses [goose](https://github.com/pressly/goose#install) migration tool).
3. Application is ready to use.
