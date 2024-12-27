package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Loop infinito!")
		break //para sair do loop
	}

	contador := 0
	for contador < 5 {
		fmt.Println(contador)
		contador++
	}

	nomes := []string{"Alice", "Bob", "Charlie"}
	for indice, nome := range nomes {
		fmt.Printf("Índice: %d, Nome: %s\n", indice, nome)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
