package main

import (
	"backend/db"
	"backend/router"
)

func main() {
	db.ConnectDatabase()

	r := router.SetUpRouter()
	r.Run(":8080")
}