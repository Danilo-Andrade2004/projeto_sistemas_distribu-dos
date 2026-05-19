package database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func Database(){
	var erro error
	
	DB, _ = sql.Open(
		"postgres", 
		"host=localhost user=postgres password=admin dbname=Questionario port=5432 sslmode=disable",
	)
	if erro != nil {
        panic(erro)
    }

	erro = DB.Ping()
	if erro != nil{
		panic(erro)
	}

	fmt.Println("Banco de dados conectado")
}

func CriarTabela(){

	fmt.Println("Criando tabelas")

	_, erro := DB.Exec(`
		CREATE TABLE IF NOT EXISTS usuarios(
			Usuario_id SERIAL PRIMARY KEY,
			curso VARCHAR(50) NOT NULL,
			turno VARCHAR(50) NOT NULL,
			sexo VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS questionario(

			id_Usuario INT PRIMARY KEY,

			sala_de_aula INT NOT NULL,
			conversar_com_colegas INT NOT NULL,
			professores INT NOT NULL,
			campus INT NOT NULL,
			emocional_semana INT NOT NULL,
			motivacao_estudos INT NOT NULL,
			ansiedade_escolar INT NOT NULL,
			voz_na_escola INT NOT NULL,
			qualidade_sono INT NOT NULL,
			bem_estar_geral INT NOT NULL,
			comentario TEXT,

			FOREIGN KEY (id_Usuario)
			REFERENCES usuarios(Usuario_id)
		);
	`)
	if erro != nil{
		panic(erro)
	}
	fmt.Println("Tabelas criadas com sucesso!")
}