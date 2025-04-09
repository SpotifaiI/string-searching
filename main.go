package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, forneça o caminho do arquivo de texto como argumento")
		return
	}

	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		return
	}
	text := string(content)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o padrão a ser buscado: ")
	pattern, _ := reader.ReadString('\n')
	pattern = strings.TrimSpace(pattern)

	if pattern == "" {
		fmt.Println("Padrão inválido. Por favor, insira algo para buscar.")
		return
	}

	startNaive := time.Now()
	naiveCount, naivePositions := NaiveSearch(text, pattern)
	naiveTime := time.Since(startNaive)

	startRK := time.Now()
	rkCount, rkPositions := RabinKarp(text, pattern)
	rkTime := time.Since(startRK)

	fmt.Println("\n=== Resultados ===")
	fmt.Printf("Padrão buscado: '%s'\n", pattern)
	fmt.Printf("Tamanho do texto: %d caracteres\n", len(text))
	fmt.Println("\nNaive Search:")
	fmt.Printf("Ocorrências: %d\n", naiveCount)
	fmt.Printf("Posições: %v\n", naivePositions)
	fmt.Printf("Tempo: %v\n", naiveTime)
	fmt.Println("\nRabin-Karp:")
	fmt.Printf("Ocorrências: %d\n", rkCount)
	fmt.Printf("Posições: %v\n", rkPositions)
	fmt.Printf("Tempo: %v\n", rkTime)

	fmt.Println("\n=== Reflexão ===")
	fmt.Printf("Os resultados são iguais? %v\n", naiveCount == rkCount)
	fmt.Printf("Qual foi mais rápido? %s\n", map[bool]string{true: "Rabin-Karp", false: "Naive Search"}[rkTime < naiveTime])
}
