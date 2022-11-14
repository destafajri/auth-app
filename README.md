# Simple Mono-Repo-app API Golang Postgres

Simple Auth-app and Fetch-app API Golang with PostgreSql and gin-gionic

## How to use
- Please clone or download this repository.
- Prepare postgres database, or use docker, you can type
```
docker-compose up
```
OR
```
docker-compose up -d
```
- add .env file to setup your database connection
```
POSTGRES_URL="user=postgres password=[your-password] host=[your-host] port=5432 dbname=postgres"
KEY_JWT="[type your secret key here]"
```
- run the golang server
```
go run main.go
```
## Postman API-Doc
# Auth-app
https://documenter.getpostman.com/view/22138766/2s8YerPCyM
# Fetch-app
https://documenter.getpostman.com/view/22138766/2s8YevqAyE

