package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"analab.com/golang/note"
	"analab.com/golang/todo"
)

func main() {
	title, content := getNoteData()
	taskContent := getTodoData()

	note, errNote := note.New(title, content)
	if errNote != nil {
		fmt.Println(errNote)
		return
	}

	task, errTask := todo.New(taskContent)
	if errTask != nil {
		fmt.Println(errTask)
		return
	}

	saveData(note)
	saveData(task)

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.Trim(text, "\n")
	text = strings.Trim(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getTodoData() string {
	content := getUserInput("Todo content:")
	return content
}

func saveData(output Outputtable) {
	output.Display()
	err := output.Save()
	if err != nil {
		fmt.Println("Could not save file")
	}
}
