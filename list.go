package main

import (
    "fmt"
)

func list(todos []Todo, project string) {
    if len(todos) == 0 {
        fmt.Println("Your TODO list is empty!")
        fmt.Println("Use \"todo add {title}\" to add new items")
    } else {
        fmt.Println("~~~~~~~~~~~~~~SPARK~TASKS~~~~~~~~~~~~~~")
    }

    for _, todo := range todos {
        if project == "" || todo.Project == project {
            checked := "[ ] "
            if todo.Completed {
                checked = "[X] "
            }
            fmt.Print(checked, todo.ID, ": ", todo.Title, " <", todo.Project, "> ", "\n")
        }
    }
}

