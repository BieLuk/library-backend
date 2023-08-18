### Simple Library Backend
This repo contains Golang backend REST API using Gin framework. 
This simple application uses Viper library for handling configuration, Gorm as sql ORM for postgresql and goose as migration tool.
</br>
Application has implemented repository for Postgresql and Mongo DB.
Application is containerised with docker-compose.

To run application:
1. Set env variables, defaults are in `app.env` file.
2. If you use Postgresql as chosen database you need to migrate Postgresql DB schema. To do that you can run script `migrate_db.sh` from `tools` catalog (the script uses [goose](https://github.com/pressly/goose#install) migration tool).
3. Run command `docker-compose up`.
4. Application is ready to use.
