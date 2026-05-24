package controller

import(
	"encoding/json"
	"net/http"
	"PROJETO_SISTEMAS_DISTRIBUIDOS/database"
)

type Usuario struct{
	Curso_Estuda int `json:"curso"`
	Turno_Estuda int `json:"turno"`
	Sexo int 		 `json:"sexo"`
}



func PostCadastrarUsuario(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil{
		http.Error(w,"Erro!, JSON inválido!", http.StatusBadRequest)
		return
	}

	var curso_nome string
	switch usuario.Curso_Estuda{
	case 1:
		curso_nome = "Informática"
	case 2:
		curso_nome = "Administração"
	case 3:
		curso_nome = "Química"
	case 4:
		curso_nome = "Adm_Subsequente"
	case 5:
		curso_nome = "TPQ"
	case 6:
		curso_nome = "TADS"
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "curso inválido",})
		return
	}

	var turno string
	switch usuario.Turno_Estuda{
	case 1:
		turno = "Matutino"
	case 2:
		turno = "Vespertino"
	case 3:
		turno = "Noturno"
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "turno inválido",})
		return
	}

	var sexo string
	switch usuario.Sexo{
	case 1:
		sexo = "Masculino"
	case 2:
		sexo = "Feminino"
	case 3:
		sexo = "Outro"
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "sexo inválido"})
		return
	}

	var usuarioID int

	err = database.DB.QueryRow(`
		INSERT INTO usuarios (curso, turno, sexo)
		VALUES ($1, $2, $3)
		RETURNING usuarioID
	`, curso_nome, turno, sexo).Scan(&usuarioID)

	if err != nil{
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "erro ao salvar no banco",
		})
	}

	type Response struct{
		UsuarioID int `json:"usuario_id"`
		Curso string `json:"curso"`
		Turno string `json:"turno"`
		Sexo string `json:"sexo"`
	}

	resp := Response{
		UsuarioID: usuarioID,
		Curso: curso_nome,
		Turno: turno,
		Sexo:  sexo,
	}

	json.NewEncoder(w).Encode(resp)
}