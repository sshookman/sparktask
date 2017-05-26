package main

import (
    "fmt"
)

func help() {
    fmt.Println("Available Commands:")
    fmt.Println("  help", "\t\t\t-", "Displays the available commands")
    fmt.Println("  list", "\t\t\t-", "Displays the list of TODOs and their status")
    fmt.Println("  add {title}", "\t\t-", "Adds a new TODO to the list")
    fmt.Println("  mark {index}", "\t\t-", "Marks or unmarks the TODO item with the given index")
    fmt.Println("  clean", "\t\t-", "Removes completed TODOs from the list")
    fmt.Println("  remove {index}", "\t-", "Removes a TODO from the list regardless of status")
}
