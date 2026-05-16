package controller

import(
	"fmt"
)

func CadastrarUsuario(){

	var curso string
	var turno string
	var numero int

	fmt.Println("Digite qual é o seu curso: ")
	fmt.Printf("1 - Administração\n2 - Química\n3 - Informática\n4 - Adm_Subsequente\n5 - TPQ\n6 - TADS\n")
	fmt.Scan(&numero)

	switch numero{
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
		fmt.Println("ERRO: Digite um número válido")
	}
	fmt.Println("--------------------------------------------------")

	fmt.Println("Digite o turno que você estuda: ")
	fmt.Printf("1 - Matutino\n2 - Vespertino\n3 - Noturno\n")
	fmt.Scan(&numero)

	switch numero {
	case 1:
		turno = "Matutino"
	case 2:
		turno = "Vespertino"
	case 3:
		turno = "Noturno"
	default:
		fmt.Println("ERRO: Digite um número válido")
	}
	fmt.Println("--------------------------------------------------")

	var sexo string
	fmt.Println("Digite o seu sexo: ")
	fmt.Println("1 - Masculino\n2 - Feminino")
	fmt.Scan(&numero)

	switch numero {
	case 1:
		sexo = "Masculino"
	case 2:
		sexo = "Feminino"
	default:
		fmt.Println("ERRO: Digite um número válido")
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("Formulário de cadastro preenchido com sucesso!")
	fmt.Println("--------------------------------------------------")
	fmt.Println("O seu curso é: ", curso)
	fmt.Println("O seu turno é: ", turno)
	fmt.Println("O seu sexo é: ", sexo)
}