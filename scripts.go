package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MenuWithBufio() {
	var reader *bufio.Reader
	for {
		reader = bufio.NewReader(os.Stdin)
		fmt.Print("Ingresa un texto: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error al leer los datos de entrada: ", err)
			return
		}

		fmt.Printf("Entrada ingresada: %v \n", input)
		if strings.TrimSpace(input) == "salida" {
			fmt.Println("Ingresaste salida")
			break
		}
	}
}

func MenuWithFmt() {
	var input string

	fmt.Print("Ingresa tu entrada por teclado: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Error en la entrada de teclado: ", err)
		return
	}

	fmt.Println("Entrada ingresada: ", input)
}
