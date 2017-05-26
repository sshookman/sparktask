package main

import (
    "fmt"
)

func clean(tasks []Task, project string) {

    var open []Task
    for _, task := range tasks {
        if project == "" || task.Project == project {
            if task.Completed {
                fmt.Println("Removed:", task.Title)
            } else {
                open = append(open, task)
            }
        }
    }

    save(open)
    list(open, project)
}
