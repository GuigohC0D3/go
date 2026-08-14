package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibirIntroducao()
	for {
		exibirOpcoes()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço esse comando!")
			os.Exit(-1)
		}
	}
}

func exibirIntroducao() {
	nome := "Guilherme"
	versao := 1.1
	fmt.Println("Olá, Sr(a).", nome)
	fmt.Printf("Sistema de Monitoramento de Sites - Versão %.1f \n", versao)
	fmt.Print("Digite uma das opções abaixo:\n")
}

func exibirOpcoes() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int {
	var liComando int
	fmt.Scan(&liComando)
	fmt.Println("O comando escolhido foi:", liComando)
	return liComando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br", "https://www.caelum.com.br"}
	fmt.Println(sites)
	
	site := "https://www.facebook.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas", resp.StatusCode)
	}
}
