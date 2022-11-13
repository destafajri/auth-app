package main

import (
	auth "github.com/destafajri/auth-app/routes"
	fetch "github.com/destafajri/fetch-app/routes"
)
func main() {
	auth.Router()
	fetch.Router()
}