package main

import (
	"github.com/RenanLourenco/go-gin.git/database"
	"github.com/RenanLourenco/go-gin.git/routes"
)

func main() {
	database.ConectarDatabase()
	routes.HandleRequests()
}