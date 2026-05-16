package controller

import(
	"fmt"
	"bufio"
	"os"
)

func Questionario(){
    var comentario string
    var numero int

    fmt.Printf("\n***AMBIENTE ESCOLAR***\n")
    fmt.Println("Como você se sente dentro da sala de aula?")
	fmt.Println("1 - BEM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BEM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    fmt.Println("Como você se sente ao conversar com colegas?")
	fmt.Println("1 - BEM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BEM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    fmt.Println("Como você se sente em relação aos professores?")
	fmt.Println("1 - BEM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BEM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    fmt.Println("Como você se sente no ambiente do campus?")
	fmt.Println("1 - BEM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BEM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    fmt.Printf("\n***SAÚDE MENTAL***\n")

    fmt.Println("Como você se sentiu emocionalmente nesta semana?")
	fmt.Println("1 - BEM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BEM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    fmt.Println("Você se sente motivado(a) para estudar?")
    fmt.Println("1 - BEM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BEM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

	fmt.Println("Você sente ansiedade durante as atividades escolares?")
    fmt.Println("1 - NÃO\n2 - MAIS OU MENOS\n3 - SIM")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("NÃO")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("SIM")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    fmt.Println("Você sente que está sendo ouvido dentro da escola?")
	fmt.Println("1 - BEM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BEM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

	fmt.Printf("\n***BEM-ESTAR***\n")

    fmt.Println("Como está sua qualidade de sono recentemente?")
	fmt.Println("1 - BOM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BOM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    fmt.Println("Como você avalia seu bem-estar geral atualmente?")
	fmt.Println("1 - BOM\n2 - MAIS OU MENOS\n3 - MAL")
    fmt.Scan(&numero)

    switch numero{
        case 1:
        fmt.Println("BOM")
        
        case 2:
        fmt.Println("MAIS OU MENOS")

        case 3:
        fmt.Println("MAL")

        default:
        fmt.Println("ERRO! digite um número válido!")
    }
    fmt.Println("**************************************************")

    leitor := bufio.NewReader(os.Stdin)
	fmt.Println("DESEJA RELATAR ALGUMA SITUAÇÃO?")
	leitor.ReadString('\n') // Limpa o buffer de entrada
	comentario, _ = leitor.ReadString('\n')
	fmt.Println("Comentário registrado:", comentario)
}