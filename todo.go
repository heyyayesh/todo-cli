package main

import (
	"errors"
	"slices"
	"time"
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
