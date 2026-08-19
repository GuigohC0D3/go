package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
		case 3:
			adicionarSite()
		case 4:
			removerSite()
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
	fmt.Println("3 - Adicionar novo Site")
	fmt.Println("4 - Remover Site")
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
		registraLog(site, false)
		return
	}

	defer resp.Body.Close()

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

		if linha != "" {
			sites = append(sites, linha)
		}

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

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

func adicionarSite() {
	var site string
	fmt.Print("Digite o link do site: ")
	fmt.Scan(&site)

	site = strings.TrimSpace(site)

	if site == "" {
		fmt.Println("Site inválido!")
		return
	}

	arquivo, err := os.OpenFile("sites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
		return
	}

	arquivo.WriteString(site + "\n")
	arquivo.Close()

	fmt.Println("Site", site, "adicionado com sucesso!")
}

func removerSite() {
	sites := lerSitesDoArquivo()

	var lista []string
	for _, s := range sites {
		if strings.TrimSpace(s) == "" {
			continue
		}
		lista = append(lista, s)
	}

	if len(lista) == 0 {
		fmt.Println("Nenhum site cadastrado!")
		return
	}

	fmt.Println("Sites cadastrados:")
	for i, s := range lista {
		fmt.Println(i+1, "-", s)
	}
	fmt.Println("0 - Cancelar")

	var escolha int
	fmt.Print("Digite o número do site que deseja remover: ")
	fmt.Scan(&escolha)

	if escolha == 0 {
		fmt.Println("Operação cancelada.")
		return
	}

	if escolha < 1 || escolha > len(lista) {
		fmt.Println("Número inválido!")
		return
	}

	site := lista[escolha-1]

	var sitesRestantes []string
	for i, s := range lista {
		if i == escolha-1 {
			continue
		}
		sitesRestantes = append(sitesRestantes, s)
	}

	conteudo := strings.Join(sitesRestantes, "\n")
	if conteudo != "" {
		conteudo = conteudo + "\n"
	}

	err := os.WriteFile("sites.txt", []byte(conteudo), 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
		return
	}

	fmt.Println("Site", site, "removido com sucesso!")
}
