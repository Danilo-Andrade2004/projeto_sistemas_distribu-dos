package controller

import(
    "encoding/json"
    "net/http"
    "PROJETO_SISTEMAS_DISTRIBUIDOS/database"
)

type Questionario struct{
    UsuarioID int `json:"usuario_id"`
    Pergunta1 int `json:"sala_de_aula"`
    Pergunta2 int `json:"conversar_com_colegas"`
    Pergunta3 int `json:"professores"`
    Pergunta4 int `json:"campus"`
    Pergunta5 int `json:"emocional_semana"`
    Pergunta6 int `json:"motivacao_estudos"`
    Pergunta7 int `json:"ansiedade_escolar"`
    Pergunta8 int `json:"voz_na_escola"`
    Pergunta9 int `json:"qualidade_sono"`
    Pergunta10 int `json:"bem_estar_geral"`
    Comentario string `json:"quer_deixar_um_comentario"`
}

func pergunta(numero int) string{
    switch numero{
        case 1:
            return "MUITO BEM"
        case 2:
            return "BEM" 
        case 3:
            return "MAIS OU MENOS" 
        case 4:
            return "MAL" 
        case 5:
            return "MUITO MAL" 
        default:
            return "ERRO! POR FAVOR DIGITE UM NÚMERO VÁLIDO!"
    }
}

func PostResponderQuestionario(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    
    var questionario Questionario
    err := json.NewDecoder(r.Body).Decode(&questionario)
    if err != nil{
        http.Error(w,"Erro!, JSON inválido!", http.StatusBadRequest)
        return
    }

    var p1, p2, p3, p4, p5, p6, p7, p8, p9, p10 string

    p1 = pergunta(questionario.Pergunta1)
    p2 = pergunta(questionario.Pergunta2)
    p3 = pergunta(questionario.Pergunta3)
    p4 = pergunta(questionario.Pergunta4)
    p5 = pergunta(questionario.Pergunta5)
    p6 = pergunta(questionario.Pergunta6)
    p7 = pergunta(questionario.Pergunta7)
    p8 = pergunta(questionario.Pergunta8)
    p9 = pergunta(questionario.Pergunta9)
    p10 = pergunta(questionario.Pergunta10)
    const mensagemErro = "ERRO! POR FAVOR DIGITE UM NÚMERO VÁLIDO!"

    if p1 == mensagemErro || p2 == mensagemErro || p3 == mensagemErro || p4 == mensagemErro ||
       p5 == mensagemErro || p6 == mensagemErro || p7 == mensagemErro || p8 == mensagemErro ||
       p9 == mensagemErro || p10 == mensagemErro{
        json.NewEncoder(w).Encode(map[string]string{
			"erro!": "Uma ou mais respostas enviadas são inválidas, use apenas 1, 2 ou 3!",
		})
        return
    }

    err = database.DB.QueryRow("SELECT usuarioID FROM usuarios WHERE usuarioID = $1",
        questionario.UsuarioID,).Scan(&questionario.UsuarioID)
    if err != nil {
        http.Error(w, "usuário não encontrado", http.StatusBadRequest)
        return
    }

    var questionarioID int

	err = database.DB.QueryRow(`
		INSERT INTO questionario(
            usuarioID, 
            sala_de_aula,
			conversar_com_colegas,
			professores,
			campus,
			emocional_semana,
			motivacao_estudos,
			ansiedade_escolar,
			voz_na_escola,
			qualidade_sono,
			bem_estar_geral, 
            comentario)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`, questionario.UsuarioID, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, questionario.Comentario).Scan(&questionarioID)

	if err != nil{
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro ao salvar no banco",
		})
        return
	}

    type Response struct{
        ID int `json:"id"`
        UsuarioID int `json:"usuario_id"`
        Pergunta1 string `json:"sala_de_aula"`
        Pergunta2 string `json:"conversar_com_colegas"`
        Pergunta3 string `json:"professores"`
        Pergunta4 string `json:"campus"`
        Pergunta5 string `json:"emocional_semana"`
        Pergunta6 string `json:"motivacao_estudos"`
        Pergunta7 string `json:"ansiedade_escolar"`
        Pergunta8 string `json:"voz_na_escola"`
        Pergunta9 string `json:"qualidade_sono"`
        Pergunta10 string `json:"bem_estar_geral"`
        Comentario string `json:"quer_deixar_um_comentario"`
    }

    resp := Response{
        ID: questionarioID,
        UsuarioID: questionario.UsuarioID,
        Pergunta1: p1,
        Pergunta2: p2,
        Pergunta3: p3,
        Pergunta4: p4,
        Pergunta5: p5,
        Pergunta6: p6,
        Pergunta7: p7,
        Pergunta8: p8,
        Pergunta9: p9,
        Pergunta10: p10,
        Comentario: questionario.Comentario,
    }

    err = json.NewEncoder(w).Encode(resp)
    if err != nil{
        http.Error(w, "Erro ao responder o JSON!", http.StatusBadRequest)
    }
}