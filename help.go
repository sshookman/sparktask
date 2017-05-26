package main

import (
    "fmt"
)

func help() {
    fmt.Println("Available Commands:")
    fmt.Println("  help", "\t\t\t-", "Displays the available commands")
    fmt.Println("  list", "\t\t\t-", "Displays the list of Tasks and their status")
    fmt.Println("  add {title}", "\t\t-", "Adds a new Task to the list")
    fmt.Println("  mark {index}", "\t\t-", "Marks or unmarks the Task item with the given index")
    fmt.Println("  clean", "\t\t-", "Removes completed Tasks from the list")
    fmt.Println("  remove {index}", "\t-", "Removes a Task from the list regardless of status")
}
