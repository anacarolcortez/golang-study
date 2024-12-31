package multiprodutores

//Várias goroutines podem enviar dados para um mesmo canal quando você deseja consolidar informações ou realizar tarefas em paralelo.

import (
	"fmt"
)

func worker(id int, ch chan string) {
	message := fmt.Sprintf("Worker %d finished task", id)
	ch <- message
}

func main() {
	ch := make(chan string)
	numWorkers := 5

	// Iniciando múltiplas goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	// Recebendo mensagens de todas as goroutines
	for i := 1; i <= numWorkers; i++ {
		fmt.Println(<-ch)
	}

	close(ch)
}

//Quando usar:
//Quando você tem várias fontes de dados que precisam ser consolidadas em um único fluxo.
//Exemplo: processamento paralelo de arquivos, envio de logs de diferentes serviços.
