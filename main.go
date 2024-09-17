package main

import (
	"log"
)

func main() {
	todos := Todos{}
	todos.add("Go to the dentist")
	todos.add("Go shopping for dress")
	err := todos.add("Send daily report")
	if err != nil {
		log.Fatal(err)
	}
	// todos.delete(1)
	todos.toggle(0)
	// todos.toggle(0)
	todos.edit(0, "Go to the dentist after this week")

	// for _, todo := range todos {
	// 	fmt.Printf("%+v\n", todo)
	// }
	todos.print()
}
