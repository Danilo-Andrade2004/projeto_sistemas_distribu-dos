


package routes

import (
	"fmt"
	"net/http"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/controller"
)

func Rotas() {
	
	http.HandleFunc("/loginadmin", controller.PostLoginAdmin)
	http.HandleFunc("/listarquestionario", controller.GetListarRespostasQuestionario)
	http.HandleFunc("/buscarquestionario", controller.GetBuscarRespostasQuestionarioPorID)
	http.HandleFunc("/cadastro", controller.PostCadastrarUsuario)
	http.HandleFunc("/questionario", controller.PostResponderQuestionario)

	
	fs := http.FileServer(http.Dir("./templates")) 
	http.Handle("/", fs)


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}