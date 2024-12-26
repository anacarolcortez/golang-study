package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Println(note.title + ": " + note.content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

// Implementação do método MarshalJSON para serialização personalizada
// Porque o Marshal precisa que os campos do struct sejam públicos,
// porém, não é desejado torná-los públicos para toda a aplicação
func (note Note) MarshalJSON() ([]byte, error) {
	type Alias struct {
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}

	return json.Marshal(&Alias{
		Title:     note.title,
		Content:   note.content,
		CreatedAt: note.createdAt,
	})
}
