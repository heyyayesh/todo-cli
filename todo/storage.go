package todo

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FilePath string
}

func NewStorage[T any] (filePath string) *Storage[T] {
	return &Storage[T]{
		FilePath: filePath,
	}
}

func (s *Storage[T]) Save(data T) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, jsonData, 0644)
}

func (s *Storage[T]) Load(into *T) error {
	byteData, err := os.ReadFile(s.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteData, into)
}

func (s *Storage[T]) Clear() error {
	return os.WriteFile(s.FilePath, []byte{}, 0644)
}
