package main

import (
	"zhihu-column-api/db"
	"zhihu-column-api/router"
)

func main() {
	db.ConnectDB()
	router.RunHTTP()
}
