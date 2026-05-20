package routes

import(
	"net/http"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/controller"
)

func Rotas(){
	http.HandleFunc("/loginadmin", controller.PostLoginAdmin)
	http.HandleFunc("/listarquestionario", controller.GetListarRespostasQuestionario)
	http.HandleFunc("/buscarquestionario", controller.GetBuscarRespostasQuestionarioPorID)
	http.HandleFunc("/cadastro", controller.PostCadastrarUsuario)
	http.HandleFunc("/questionario", controller.PostResponderQuestionario)
	http.ListenAndServe(":8080", nil)
}