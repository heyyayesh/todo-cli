package cmd

import (
	"www.github.com/heyyayesh/todo-cli/todo"
)

func LoadData(todos *todo.Todos) error {
	storage := todo.NewStorage[todo.Todos]("todos.json")
	return storage.Load(todos)
}

func SaveData(todos todo.Todos) error {
	storage := todo.NewStorage[todo.Todos]("todos.json")
	return storage.Save(todos)
}

func ClearData() error {
	storage := todo.NewStorage[todo.Todos]("todos.json")
	return storage.Clear()
}
