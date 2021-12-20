package main

import "fmt"

func main() {

	var num1, num2 float64
	var operacao string

	fmt.Println("Selecione a operação (+ - * /)")
	fmt.Scanln(&operacao)

	fmt.Println("Digite o primeiro valor")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo valor")
	fmt.Scanln(&num2)

	switch operacao {

	case "+":
		fmt.Printf("%0.2f %s %0.2f = %0.2f\n", num1, operacao, num2, num1+num2)

	case "-":
		fmt.Printf("%0.2f %s %0.2f = %0.2f\n", num1, operacao, num2, num1-num2)

	case "*":
		fmt.Printf("%0.2f %s %0.2f = %0.2f\n", num1, operacao, num2, num1*num2)

	case "/":
		fmt.Printf("%0.2f %s %0.2f = %0.2f\n/", num1, operacao, num2, num1/num2)

	}

}
