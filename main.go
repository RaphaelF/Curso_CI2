package main

import (
	"./database"
	"github.com/guilhermeonrails/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
