package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
		time.Sleep(1 * time.Second)
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
}

func readFile() []string {
	// arq, err := os.Open("sites.txt")

	arq, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Print("ERROR", err)
	}
	fmt.Println(string(arq))
	return []string{""}
}
