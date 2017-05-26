package main

func add(todos []Todo, title, project string) {
    index := getCurrentIndex(todos) + 1
    todo := Todo{index, title, false, project}
    todos = append(todos, todo)

    save(todos)
    list(todos, project)
}

func getCurrentIndex(todos []Todo) (index int) {

    index = 0
    for _, todo := range todos {
        if index < todo.ID {
            index = todo.ID
        }
    }

    return
}
