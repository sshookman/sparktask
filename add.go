package main

func add(tasks []Task, title, project string) {
    index := getCurrentIndex(tasks) + 1
    task := Task{index, title, false, project}
    tasks = append(tasks, task)

    save(tasks)
    list(tasks, project)
}

func getCurrentIndex(tasks []Task) (index int) {

    index = 0
    for _, task := range tasks {
        if index < task.ID {
            index = task.ID
        }
    }

    return
}
