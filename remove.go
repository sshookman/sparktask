package main

import (
    "fmt"
    "strconv"
)

func remove(tasks []Task, id string) {
    index, _ := strconv.Atoi(id)

    var remaining []Task
    for _, task := range tasks {
        if task.ID == index {
            fmt.Println("Removed:", task.Title)
        } else {
            remaining = append(remaining, task)
        }
    }

    save(remaining)
    list(remaining, "")
}
