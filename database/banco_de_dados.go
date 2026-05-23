package database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func Database(){
	var erro error
	
	DB, erro = sql.Open(
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

		CREATE TABLE IF NOT EXISTS admin(
			email VARCHAR(100) NOT NULL UNIQUE,
			senha VARCHAR(100) NOT NULL
		);
		INSERT INTO admin (email, senha)
		VALUES ('servicosocial@gmail.com', '123456')
		ON CONFLICT (email) DO NOTHING;

		CREATE TABLE IF NOT EXISTS usuarios(
			usuarioID SERIAL PRIMARY KEY,
			curso VARCHAR(50) NOT NULL,
			turno VARCHAR(50) NOT NULL,
			sexo VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS questionario(
			id SERIAL PRIMARY KEY,
			usuarioID INT NOT NULL,

			sala_de_aula VARCHAR(50) NOT NULL,
			conversar_com_colegas VARCHAR(50) NOT NULL,
			professores VARCHAR(50) NOT NULL,
			campus VARCHAR(50) NOT NULL,
			emocional_semana VARCHAR(50) NOT NULL,
			motivacao_estudos VARCHAR(50) NOT NULL,
			ansiedade_escolar VARCHAR(50) NOT NULL,
			voz_na_escola VARCHAR(50) NOT NULL,
			qualidade_sono VARCHAR(50) NOT NULL,
			bem_estar_geral VARCHAR(50) NOT NULL,
			comentario TEXT,

			FOREIGN KEY (usuarioID)
			REFERENCES usuarios(usuarioID)
		);
	`)
	if erro != nil{
		panic(erro)
	}
	fmt.Println("Tabelas criadas com sucesso!")
}