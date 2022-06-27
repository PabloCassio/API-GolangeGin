package main

import (
	"github.com/PabloCassio/api-go-gin/database"
	"github.com/PabloCassio/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
