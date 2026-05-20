package controller

import (
	"PROJETO_SISTEMAS_DISTRIBUIDOS/database"
	"encoding/json"
	"net/http"
)

type Admin struct{
	Email string `json:"email"`
	Senha string `json:"senha"`
}



func PostLoginAdmin(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var email string

	var cadastroAdmin Admin
	err := json.NewDecoder(r.Body).Decode(&cadastroAdmin)
	if err != nil{
		http.Error(w,"Erro!, JSON inválido!", http.StatusBadRequest)
		return
	}

	err = database.DB.QueryRow(
		"SELECT email FROM admin WHERE email = $1 AND senha = $2",
		cadastroAdmin.Email,
		cadastroAdmin.Senha,
	).Scan(&email)
	if err != nil{
		http.Error(w,"Erro! Email e/ou Senha inválidos!", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "login aprovado",
	})
}



func GetListarRespostasQuestionario(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	type Resposta struct{
		ID int `json:"id"`
		UsuarioID int `json:"usuario_id"`
		SalaDeAula string `json:"sala_de_aula"`
		ConversarComColegas string `json:"conversar_com_colegas"`
		Professores string `json:"professores"`
		Campus string `json:"campus"`
		EmocionalSemana string `json:"emocional_semana"`
		MotivacaoEstudos string `json:"motivacao_estudos"`
		AnsiedadeEscolar string `json:"ansiedade_escolar"`
		VozNaEscola string `json:"voz_na_escola"`
		QualidadeSono string `json:"qualidade_sono"`
		BemEstarGeral string `json:"bem_estar_geral"`
		Comentario string `json:"comentario"`
	}

	rows, err := database.DB.Query(`
		SELECT 
		id,
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
		comentario
		FROM questionario
	`)

	if err != nil{
		http.Error(w, "erro ao buscar questionários", 500)
		return
	}

	defer rows.Close()

	var respostas []Resposta

	for rows.Next(){
		var resposta Resposta
		err := rows.Scan(
			&resposta.ID,
			&resposta.UsuarioID,
			&resposta.SalaDeAula,
			&resposta.ConversarComColegas,
			&resposta.Professores,
			&resposta.Campus,
			&resposta.EmocionalSemana,
			&resposta.MotivacaoEstudos,
			&resposta.AnsiedadeEscolar,
			&resposta.VozNaEscola,
			&resposta.QualidadeSono,
			&resposta.BemEstarGeral,
			&resposta.Comentario,
		)
		if err != nil{
			http.Error(w, "erro ao ler linha", 500)
			return
		}
		respostas = append(respostas, resposta)
	}

	if err = rows.Err(); err != nil{
		http.Error(w, "erro no loop", 500)
		return
	}

	json.NewEncoder(w).Encode(respostas)
}



func GetBuscarRespostasQuestionarioPorID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")

	type Resposta struct{
		ID int `json:"id"`
		UsuarioID int `json:"usuario_id"`
		SalaDeAula string `json:"sala_de_aula"`
		ConversarComColegas string `json:"conversar_com_colegas"`
		Professores string `json:"professores"`
		Campus string `json:"campus"`
		EmocionalSemana string `json:"emocional_semana"`
		MotivacaoEstudos string `json:"motivacao_estudos"`
		AnsiedadeEscolar string `json:"ansiedade_escolar"`
		VozNaEscola string `json:"voz_na_escola"`
		QualidadeSono string `json:"qualidade_sono"`
		BemEstarGeral string `json:"bem_estar_geral"`
		Comentario string `json:"comentario"`
	}

	var err error
	var resposta Resposta

	err = database.DB.QueryRow(`
		SELECT
		id,
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
		comentario
		FROM questionario
		WHERE id = $1
	`, id).Scan(
		&resposta.ID,
		&resposta.UsuarioID,
		&resposta.SalaDeAula,
		&resposta.ConversarComColegas,
		&resposta.Professores,
		&resposta.Campus,
		&resposta.EmocionalSemana,
		&resposta.MotivacaoEstudos,
		&resposta.AnsiedadeEscolar,
		&resposta.VozNaEscola,
		&resposta.QualidadeSono,
		&resposta.BemEstarGeral,
		&resposta.Comentario)

	if err != nil{
		http.Error(w, "questionário não encontrado", 404)
		return
	}

	json.NewEncoder(w).Encode(resposta)
}