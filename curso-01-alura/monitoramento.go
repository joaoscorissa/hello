package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const green = "\033[32m"
const red = "\033[31m"
const resetColor = "\033[0m"

func main() {
	for {
		showMenu()
		comando := readCmd()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs...")
			printLogs()
		case 0:
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido :(")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("")
}

func readCmd() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando Monitoramento...")
	sites := readFile()

	for i := 0; i < 10; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(5 * time.Second)
	}
}

func testaSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Print("ERROR", err)
	}

	if response.StatusCode == 200 {
		fmt.Println(green+site, "STATUS:", response.Status, "- METHOD:", response.Request.Method+resetColor)
	} else {
		fmt.Println(red+"ERROR -", response.Status, "- METHOD:", response.Request.Method+resetColor)
	}
	registraLog(site, response)
}

func readFile() []string {
	arq, err := os.Open("sites.txt")
	var sites []string
	if err != nil {
		fmt.Print("ERROR", err)
	}

	leitor := bufio.NewReader(arq)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, string(linha))

		if err == io.EOF {
			break
		}
	}
	arq.Close()

	return sites
}

func registraLog(site string, response *http.Response) {
	arq, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {

	}

	dt := "[" + time.Now().Format("02/01/2006 15:04:05") + "]"
	if response.StatusCode == 200 {
		arq.WriteString(dt + " - \"" + site + "\" STATUS: " + response.Status + " - METHOD: " + response.Request.Method + "\n")
	} else {
		arq.WriteString(dt + " - \"" + site + "\" ERROR: " + response.Status + " - METHOD: " + response.Request.Method + "\n")
	}

	arq.Close()
}

func printLogs() {
	arq, err := ioutil.ReadFile("log.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arq))
}
