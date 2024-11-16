package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xsamael/go-cli-task-crud/internal/tasks"
)

func main() {
	var tasksList []tasks.Task

	err := tasks.ReadJson("task.json", &tasksList)

	if err != nil {
		fmt.Println("Error leyendo tareas:", err)
		return
	}

	if len(os.Args) < 2 {
		tasks.Options()
		return
	}

	switch os.Args[1] {
	//------------------------------------------------------//
	//Add task
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Uso: add <Nueva descripción>")
			return
		}
		description := os.Args[2]

		tasksList = tasks.AddTask(tasksList, description)

		err := tasks.SaveTask("task.json", tasksList)
		if err != nil {
			fmt.Println("Error guardando tareas:", err)
			return
		}
		fmt.Println("Tarea agregada con éxito.")
		//------------------------------------------------------//
		//List all task
	case "list":

		if len(os.Args) == 2 {
			tasks.ListTask(tasksList)
			return
		}
		if os.Args[2] == "done" {
			status := os.Args[2]
			tasks.ListTaskDone(tasksList, status)
			return
		}
		if os.Args[2] == "todo" {
			status := os.Args[2]
			tasks.ListTaskTodo(tasksList, status)
			return
		}
		if os.Args[2] == "in-progress" {
			status := os.Args[2]
			tasks.ListInProgress(tasksList, status)
			return
		}

		//------------------------------------------------------//
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Uso: update <ID> <Nueva descripción>")
			return
		}
		id := os.Args[2]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error: el ID debe ser un número válido.")
			return
		}
		description := os.Args[3]

		update := tasks.UpdateTask(tasksList, idInt, description)
		if !update {
			fmt.Println("Error: tarea no encontrada con el ID especificado.")
			return
		}
		err = tasks.SaveTask("task.json", tasksList)
		if err != nil {
			fmt.Println("Error al guardar las tareas:", err)
			return
		}

		fmt.Println("Tarea actualizada con éxito.")
		//------------------------------------------------------//
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Uso: delete <ID>")
			return
		}
		id := os.Args[2]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error: el ID debe ser un número válido.")
			return
		}
		delete := tasks.DeleteTask(&tasksList, idInt)
		if !delete {
			fmt.Println("Error: tarea no encontrada con el ID especificado.")
			return
		}

		err = tasks.SaveTask("task.json", tasksList)
		if err != nil {
			fmt.Println("Error al guardar las tareas:", err)
			return
		}
		fmt.Println("Tarea eliminada con éxito.")
		//------------------------------------------------------//
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Uso de mark-in-progress <ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ingresa un numero")
			return
		}
		markProgress := tasks.MarkTaskInProgress(tasksList, id)
		if !markProgress {
			fmt.Println("Tarea no encontrada")
		}

		err = tasks.SaveTask("task.json", tasksList)
		if err != nil {
			fmt.Println("Error al guardar las tareas:", err)
			return
		}
		fmt.Println("Tarea Marcada con exito")

		//------------------------------------------------------//
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Uso de mark-done <ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Ingresa un numero")
			return
		}
		done := tasks.MarkTaskAsDone(tasksList, id)
		if !done {
			fmt.Println("Tarea no encontrada")
		}

		err = tasks.SaveTask("task.json", tasksList)
		if err != nil {
			fmt.Println("Error al guardar las tareas:", err)
			return
		}
		fmt.Println("Tarea realizada")

		//------------------------------------------------------//
	default:
		fmt.Println("Comando no reconocido.")
	}
}
