# Simple Fetch-app API with Golang

Simple Fetch-app API with Golang and gin-gionic

## How to use
- Please clone or download this repository.
- Prepare postgres database, or use docker, you can type
```
docker-build --tag fetch:fetch .
```
- add .env file to setup your database connection
```
POSTGRES_URL="user=postgres password=[your-password] host=[your-host] port=5432 dbname=postgres"
KEY_JWT="[type your secret key here]"
URL="[type URL source here]"
```
- add respone helper to help json response

Sample helper
```
package helper

type ResponseFetch struct {
	ID     				string 	`json:"uuid"`
	Komoditas 			string 	`json:"komoditas"`
	Area_provinsi		string 	`json:"area_provinsi"`
	Area_kota			string 	`json:"area_kota"`
	Size				string	`json:"size"`
	Price				string	`json:"price"`
	Tgl_parsed			string	`json:"tgl_parsed"`
	Timestamp			string	`json:"timestamp"`
}
```
- run the golang server
```
go run main.go
```
## Postman API-Doc
https://documenter.getpostman.com/view/22138766/2s8YevqAyE
