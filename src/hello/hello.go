package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)

const monitoramentos = 5
const delay = 10

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
			imprimirLogs()
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

	sites := lerSitesDoArquivo()

	fmt.Println(sites)

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string
	// arquivo, err := os.Open("sites.txt")

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimirLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	fmt.Println(string(arquivo))
}