package main

import (
	"PROJETO_SISTEMAS_DISTRIBUIDOS/database"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/routes"
)

func main(){
	database.Database()
	// database.CriarTabela()
	// defer database.DB.Close()
	routes.Rotas()
}