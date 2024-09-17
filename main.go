package main

func main() {
	storage := NewStorage[Todos]("todos.json")
	todos := Todos{}
	storage.Load(&todos)

	todos.toggle(0)
	todos.delete(7)

	todos.print()
	storage.Save(todos)
}
