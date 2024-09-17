package main

import (
	"errors"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	CreatedAt time.Time
	CompletedAt *time.Time
	IsCompleted bool
}

type Todos []Todo

func (todos *Todos) add(title string) error {
	if title == "" {
		return errors.New("invalid todo title")
	}

	newTodo := Todo{
		Title: title,
		CreatedAt: time.Now(),
		CompletedAt: nil,
		IsCompleted: false,
	}

	*todos = append(*todos, newTodo)
	return nil
}

func (todos *Todos) delete(id int) error {
	if err := todos.validateIndex(id); err != nil {
		return err
	}

	t := *todos
	*todos = slices.Concat(t[:id], t[id + 1:])

	return nil
}

func (todos *Todos) toggle(id int) error {
	if err := todos.validateIndex(id); err != nil {
		return err
	}

	t := *todos
	if !(t[id].IsCompleted) {
		t[id].IsCompleted = true
		completionTime := time.Now()
		t[id].CompletedAt = &completionTime
	} else {
		t[id].IsCompleted = false
		t[id].CompletedAt = nil
	}

	return nil
}

func (todos *Todos) edit(id int, newTitle string) error {
	if err := todos.validateIndex(id); err != nil {
		return err
	}

	t := *todos
	t[id].Title = newTitle

	return nil
}

func (todos *Todos) validateIndex(idx int) error {
	if idx < 0 || idx >= len(*todos) {
		return errors.New("invalid id")
	}

	return nil
}

func (todos *Todos) print() {
	t := table.New(os.Stdout)
	t.SetRowLines(false)

	t.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")

	for idx, todo := range *todos {
		isCompleted := "❌"
		if todo.IsCompleted {
			isCompleted = "✅"
		}
		completedAt := "--"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		}

		t.AddRow(
			strconv.Itoa(idx + 1), 
			todo.Title, 
			isCompleted, 
			todo.CreatedAt.Format(time.RFC1123), 
			completedAt,
		)
	}

	t.Render()
}
