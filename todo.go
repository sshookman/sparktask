package main

import (
    "fmt"
)

func main() {
    args := parseArgs()
    project := args["p"]

    command := parseArg(1, "list")
    param1 := parseArg(2, "")
    param2 := parseArg(3, "")

    todos := load()

    switch command {
        case "help": help()
        case "list": list(todos, project)
        case "add": add(todos, param1, param2)
        case "mark": mark(todos, param1)
        case "clean": clean(todos, project)
        case "remove": remove(todos, param1)
        default: fmt.Println("Invalid Command", "-", "Use \"todo help\" to get a list of commands")
    }
}
