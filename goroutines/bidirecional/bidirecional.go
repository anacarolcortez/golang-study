package bidirecional

//Várias goroutines podem enviar e receber do mesmo canal para realizar trocas de mensagens.

import (
	"fmt"
)

func worker(id int, ch chan string) {
	ch <- fmt.Sprintf("Worker %d sending message", id)
	msg := <-ch
	fmt.Printf("Worker %d received: %s\n", id, msg)
}

func main() {
	ch := make(chan string)
	numWorkers := 2

	// Iniciando goroutines que enviam e recebem do mesmo canal
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	// Gerenciando mensagens no canal
	for i := 1; i <= numWorkers; i++ {
		msg := <-ch
		fmt.Println("Main received:", msg)
		ch <- "Acknowledged by main"
	}
}

//Quando usar:
//Quando é necessário trocar mensagens entre diferentes goroutines de forma sincronizada.
//Exemplo: protocolos de handshakes, sistemas de requisição-resposta.
