package main

import (
    "strconv"
)

func mark(todos []Todo, id string) {
    index, _ := strconv.Atoi(id)

    for i, todo := range todos {
        if todo.ID == index {
            todos[i].Completed = !todo.Completed
        }
    }

    save(todos)
    list(todos, "")
}
