package main

import "fmt"

func main() {
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Informe o valor investido (em dólares): ")
	if _, err := fmt.Scan(&investmentAmount); err != nil {
		fmt.Println("Erro: Por favor, insira um número válido para o valor investido.")
		return
	}

	fmt.Print("Informe a taxa de retorno anual (em %): ")
	if _, err := fmt.Scan(&expectedReturnRate); err != nil {
		fmt.Println("Erro: Por favor, insira um número válido para a taxa de retorno.")
		return
	}

	fmt.Print("Informe por quantos anos o valor renderá: ")
	if _, err := fmt.Scan(&years); err != nil {
		fmt.Println("Erro: Por favor, insira um número válido para os anos.")
		return
	}

	Calculate(investmentAmount, expectedReturnRate, years)

	//EXEC2
	revenue := getUserInput("Informe a receita: ")
	expenses := getUserInput("Informe as despesas: ")
	taxRate := getUserInput("Informe a alíquota de impostos: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("O Ebitda é de: US$ %.1f \n", ebt)
	fmt.Printf("O lucro é de: US$ %.1f \n", profit)
	fmt.Printf("Ebitda/lucro: %.3f \n", ratio)
}
