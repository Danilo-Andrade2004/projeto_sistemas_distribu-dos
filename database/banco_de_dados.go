package database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func Database(){
	
	DB, _ = sql.Open(
		"postgres", 
		"host=localhost user=postgres password=admin dbname=Questionario port=5432 sslmode=disable",
	)

	erro := DB.Ping()
	if erro != nil{
		panic(erro)
	}

	fmt.Println("Banco de dados conectado")
}

func CriarTabela(){

	fmt.Println("Criando tabelas")

	_, erro := DB.Exec(`
		CREATE TABLE IF NOT EXISTS usuarios(
			id_Usuario SERIAL PRIMARY KEY,
			curso VARCHAR(50) NOT NULL,
			turno VARCHAR(50) NOT NULL,
			sexo VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS questionarios(
			id SERIAL PRIMARY KEY,
			pergunta1 INT NOT NULL,
			pergunta2 INT NOT NULL,
			pergunta3 INT NOT NULL,
			pergunta4 INT NOT NULL,
			pergunta5 INT NOT NULL,
			pergunta6 INT NOT NULL,
			pergunta7 INT NOT NULL,
			pergunta8 INT NOT NULL,
			pergunta9 INT NOT NULL,
			pergunta10 INT NOT NULL,

			comentario TEXT,

			id_Usuario INT NOT NULL,

			FOREIGN KEY (id_Usuario)
			REFERENCES usuarios(id_Usuario)
		);
	`)
	if erro != nil{
		panic(erro)
	}
	fmt.Println("Tabelas criadas com sucesso!")
}