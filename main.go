package main

import (
	"PROJETO_SISTEMAS_DISTRIBUIDOS/database"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/routes"
)

func main(){
	database.Database()
	database.CriarTabela()//para criar as tabelas no banco de dados e nao ficar um banco vazio;
	routes.Rotas()
}