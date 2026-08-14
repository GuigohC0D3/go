package main

import "fmt"

func main() {
	nome := "Guilherme"
	versao := 1.1
	fmt.Println("Olá, Sr(a).", nome)

	fmt.Printf("Sistema de Monitoramento de Sites - Versão %.1f \n", versao)
	fmt.Print("Digite uma das opções abaixo:\n")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)
}
