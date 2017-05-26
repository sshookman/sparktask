package main

import (
    "os"
    "fmt"
    "strings"
    "io/ioutil"
    "encoding/json"
)

type Task struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
    Project   string `json:"project"`
}


func load() (todos []Task) {

    raw, err := ioutil.ReadFile("/root/.todos.json")
    if err != nil && strings.Contains(err.Error(), "no such file") {

        raw, err = json.Marshal("[]")
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }

        err = ioutil.WriteFile("/root/.todos.json", raw, 0644)
    }

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    json.Unmarshal(raw, &todos)
    return
}

func save(todos []Task) {
    raw, err := json.Marshal(todos)

    ioutil.WriteFile("/root/.todos.json", raw, 0644)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}
