package main

import (
    "fmt"
)

func list(tasks []Task, project string) {
    if len(tasks) == 0 {
        fmt.Println("Your TODO list is empty!")
        fmt.Println("Use \"task add {title}\" to add new items")
    } else {
        fmt.Println("~~~~~~~~~~~~~~SPARK~TASKS~~~~~~~~~~~~~~")
    }

    for _, task := range tasks {
        if project == "" || task.Project == project {
            checked := "[ ] "
            if task.Completed {
                checked = "[X] "
            }
            fmt.Print(checked, task.ID, ": ", task.Title, " <", task.Project, "> ", "\n")
        }
    }
}

