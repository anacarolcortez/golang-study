package main

import "fmt"

func main() {

	//Um array em Go tem um tamanho fixo declarado no momento da criação. Por exemplo:
	prices := [6]float64{799.0, 125.90, 5.9, 366.6, 69.9, 499.0}
	for index := range prices {
		fmt.Printf("Preço: %.2f\n", prices[index])
	}

	temperature := [5]float64{}
	temperature[0] = 23.5
	temperature[1] = 25.0
	temperature[2] = 26.5
	temperature[3] = 27.0
	temperature[4] = 28.5

	for index := range temperature {
		fmt.Printf("Temperatura: %.2f\n", temperature[index])
	}

	//Um slice é uma estrutura mais flexível construída sobre arrays.
	//Ele pode crescer ou encolher dinamicamente e não exige um tamanho fixo no momento da declaração.
	var names []string
	names = append(names, "Maria")
	names = append(names, "João")
	names = append(names, "José")
	names = append(names, "Ana")

	for index, name := range names {
		fmt.Printf("Índice: %d, Nome: %s\n", index, name)
	}

	fmt.Println("Tamanho do slice de nomes:", len(names))
	fmt.Println("Capacidade do slice de nomes:", cap(names))
	fmt.Println(names[2])

	//Slice pode ser um recorte de um array
	featurePrices := prices[0:4] // não inclui o último índice fornecido
	highlightedPrices := featurePrices[:2]
	fmt.Println(featurePrices)
	fmt.Println(highlightedPrices)

	//Cuidado! Um slice, enquanto um pedaço de um array, pode alterar valores no array original
	// O slice é uma referência a um trecho de um array, e não uma cópia
	testPrices := prices[0:4]
	testPrices[2] = 55.9
	fmt.Println(prices)
}
