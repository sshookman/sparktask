package main

import (
    "fmt"
    "strconv"
)

func remove(todos []Todo, id string) {
    index, _ := strconv.Atoi(id)

    var remaining []Todo
    for _, todo := range todos {
        if todo.ID == index {
            fmt.Println("Removed:", todo.Title)
        } else {
            remaining = append(remaining, todo)
        }
    }

    save(remaining)
    list(remaining, "")
}
