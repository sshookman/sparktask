package main

import (
    "strconv"
)

func mark(tasks []Task, id string) {
    index, _ := strconv.Atoi(id)

    for i, task := range tasks {
        if task.ID == index {
            tasks[i].Completed = !task.Completed
        }
    }

    save(tasks)
    list(tasks, "")
}
