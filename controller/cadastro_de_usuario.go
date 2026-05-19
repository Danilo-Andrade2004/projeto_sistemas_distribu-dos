package controller

import(
	"encoding/json"
	"net/http"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/database"
)

type Usuario struct{
	Curso string `json:"curso"`
	Turno string `json:"turno"`
	Sexo string `json:"sexo"`
}

func CadastrarUsuario(
	w http.ResponseWriter,
	r *http.Request,){

		w.Header().Set("Content-Type", "application/json")

		var input struct {
		Curso int `json:"curso"`
		Turno int `json:"turno"`
		Sexo  int `json:"sexo"`
	    }

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": "JSON inválido",
			})
			return
		}

		var curso string
		switch input.Curso{
		case 1:
			curso = "Administração"
		case 2:
			curso = "Química"
		case 3:
			curso = "Informática"
		case 4:
			curso = "Adm_Subsequente"
		case 5:
			curso = "TPQ"
		case 6:
			curso = "TADS"
		default:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": "curso inválido",
			})
		return
		}

		var turno string
		switch input.Turno{
			case 1:
				turno = "Matutino"
			case 2:
				turno = "Vespertino"
			case 3:
				turno = "Noturno"
			default:
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"erro": "turno inválido",
				})
				return
		}

		var sexo string
		switch input.Sexo{
			case 1:
				sexo = "Masculino"
			case 2:
				sexo = "Feminino"
			default:
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"erro": "sexo inválido",
				})
				return
		}

		query := `INSERT INTO usuarios (curso, turno, sexo) VALUES ($1, $2, $3) RETURNING Usuario_id`
	
		var usuarioID int
	
		
		err = database.DB.QueryRow(query, curso, turno, sexo).Scan(&usuarioID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": "Erro ao salvar no banco de dados: " + err.Error(),
			})
			return
		}

		resposta := Usuario{
			Curso: curso,
			Turno: turno,
			Sexo:  sexo,
		}

    	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		
    	mensagem := "🎉 Cadastro Realizado com Sucesso!\n" +
    	            "---------------------------------\n" +
    	            "📚 Curso: " + resposta.Curso + "\n" +
    	            "⏰ Turno: " + resposta.Turno + "\n" +
    	            "👤 Sexo:  " + resposta.Sexo + "\n" +
    	            "---------------------------------"
    	w.Write([]byte(mensagem))

		// json.NewEncoder(w).Encode(resposta)
}