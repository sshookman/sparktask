package main

import (
    "fmt"
)

func clean(todos []Todo, project string) {

    var open []Todo
    for _, todo := range todos {
        if project == "" || todo.Project == project {
            if todo.Completed {
                fmt.Println("Removed:", todo.Title)
            } else {
                open = append(open, todo)
            }
        }
    }

    save(open)
    list(open, project)
}
