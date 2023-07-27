#!/bin/bash

source app.env
GOOSE_DRIVER="postgres";
GOOSE_DB_STRING="host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable";

goose -dir db/migrate $GOOSE_DRIVER "$GOOSE_DB_STRING" up
