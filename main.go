package main

import (
	"PROJETO_SISTEMAS_DISTRIBUIDOS/database"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/routes"
)

func main(){
	database.Database()
	routes.Rotas()
}