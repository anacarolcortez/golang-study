package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	content   string
	createdAt time.Time
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content cannot be empty")
	}

	return Todo{
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.content)
}

func (todo Todo) Save() error {
	fileName := "todo_" + todo.createdAt.Format("2006-01-02") + ".json"

	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

// Implementação do método MarshalJSON para serialização personalizada
// Porque o Marshal precisa que os campos do struct sejam públicos,
// porém, não é desejado torná-los públicos para toda a aplicação
func (todo Todo) MarshalJSON() ([]byte, error) {
	type Alias struct {
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}

	return json.Marshal(&Alias{
		Content:   todo.content,
		CreatedAt: todo.createdAt,
	})
}
