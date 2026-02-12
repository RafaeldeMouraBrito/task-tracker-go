package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func addTask(task string) {
	dataRead, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	var lista []Task
	err = json.Unmarshal(dataRead, &lista)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	maxId := 0
	for _, t := range lista {
		if t.ID > maxId {
			maxId = t.ID
		}
	}
	fmt.Printf("O maior ID encontrado foi: %d\n", maxId)
	proximoId := maxId + 1

	newTask := Task{
		ID:          proximoId,
		Description: task,
		Status:      "todo",
	}

	lista = append(lista, newTask)
	data, err := json.Marshal(lista)
	if err != nil {
		fmt.Println("Erro ao transformar o arquivo:", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, especifique um comando: add, list, update, delete, mark-in-progress, mark-complete")
		return
	}
	var typeOperation string = os.Args[1]

	switch typeOperation {
	case "add":
		addTask(os.Args[2])
	case "list":
		fmt.Println("list!")
	case "update":
		fmt.Println("update!")
	case "delete":
		fmt.Println("delete!")
	case "mark-in-progress":
		fmt.Println("mark-in-progress!")
	case "mark-complete":
		fmt.Println("mark-complete!")
	default:
		fmt.Println("Unknown command")
	}
}
