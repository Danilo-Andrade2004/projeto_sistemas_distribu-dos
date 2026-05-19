package routes

import(
	"net/http"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/controller"
)

func Rotas(){
	http.HandleFunc("/cadastro", controller.CadastrarUsuario)
	http.HandleFunc("/questionario", controller.Questionario)

	http.ListenAndServe(":8080", nil)
}