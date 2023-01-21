package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Lectura de la entrada del usuario
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()

	fmt.Println(operacion)
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])

	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])

	fmt.Println(operador1 + operador2)

}
