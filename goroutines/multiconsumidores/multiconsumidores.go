package multiconsumidores

//Múltiplas goroutines podem ler de um mesmo canal para distribuir a carga de trabalho (padrão de "worker pool").

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for task := range ch {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(500 * time.Millisecond) // Simula tempo de processamento
	}
}

func main() {
	ch := make(chan int)
	numWorkers := 3
	numTasks := 10

	// Iniciando múltiplos workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	// Enviando tarefas para o canal
	for i := 1; i <= numTasks; i++ {
		ch <- i
	}

	close(ch) // Fecha o canal quando todas as tarefas forem enviadas

	// Esperando os workers terminarem (simplesmente aguarda)
	time.Sleep(5 * time.Second)
}

//Quando usar:
//Quando você deseja dividir o trabalho entre múltiplos consumidores para melhorar a eficiência.
//Exemplo: filas de processamento, renderização paralela, crawling na web.
