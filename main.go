package main

import (
	// "PROJETO_SISTEMAS_DISTRIBUIDOS/controller"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/database"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/routes"
)

func main(){
	database.Database()
	defer database.DB.Close()
	routes.Rotas()
	// controller.CadastrarUsuario()
	// controller.Questionario()
}