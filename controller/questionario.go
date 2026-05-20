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
    Comentario string `json:"quer_deixar_um_comentario?"`
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

    switch questionario.Pergunta1{
        case 1:
        p1 = "Bem"
        case 2:
        p1 = "Mais ou menos"
        case 3:
        p1 = "Mal"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 1",})
        return
    }

    switch questionario.Pergunta2{
        case 1:
        p2 = "Bem"
        case 2:
        p2 = "Mais ou menos"
        case 3:
        p2 = "Mal"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 2",})
        return
    }

    switch questionario.Pergunta3{
        case 1:
        p3 = "Bem"
        case 2:
        p3 = "Mais ou menos"
        case 3:
        p3 = "Mal"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 3",})
        return
    }

    switch questionario.Pergunta4{
        case 1:
        p4 = "Bem"
        case 2:
        p4 = "Mais ou menos"
        case 3:
        p4 = "Mal"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 4",})
        return
    }

    switch questionario.Pergunta5{
        case 1:
        p5 = "Bom"
        case 2:
        p5 = "Mais ou menos"
        case 3:
        p5 = "Ruim"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 5",})
        return
    }

    switch questionario.Pergunta6{
        case 1:
        p6 = "Sim"
        case 2:
        p6 = "Mais ou menos"
        case 3:
        p6 = "Não"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 6",})
        return
    }

    switch questionario.Pergunta7{
        case 1:
        p7 = "Não"
        case 2:
        p7 = "Mais ou menos"
        case 3:
        p7 = "Sim"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 7",})
        return
    }

    switch questionario.Pergunta8{
        case 1:
        p8 = "Sim"
        case 2:
        p8 = "Mais ou menos"
        case 3:
        p8 = "Não"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 8",})
        return
    }

    switch questionario.Pergunta9{
        case 1:
        p9 = "Bom"
        case 2:
        p9 = "Mais ou menos"
        case 3:
        p9 = "Ruim"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 9",})
        return
    }

    switch questionario.Pergunta10{
        case 1:
        p10 = "Bem"
        case 2:
        p10 = "Mais ou menos"
        case 3:
        p10 = "Mal"
        default:
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "erro": "resposta inválida para pergunta 10",})
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
        Comentario string `json:"quer_deixar_um_comentario?"`
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