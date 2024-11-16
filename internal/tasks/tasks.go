package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/xsamael/go-cli-task-crud/internal/model"
)

type Task = model.Task

func Options() {
	fmt.Println("Uso:")
	fmt.Println("  add <descripcion>               - Agrega una nueva tarea")
	fmt.Println("  list                           - Muestra todas las tareas")
	fmt.Println("  list-done                      - Lista todas las tareas marcadas como 'done'")
	fmt.Println("  list-todo                      - Lista todas las tareas marcadas como 'todo'")
	fmt.Println("  list-in-progress               - Lista todas las tareas marcadas como 'in-progress'")
	fmt.Println("  update <id> <descripcion>      - Actualiza una tarea por su ID")
	fmt.Println("  delete <id>                    - Elimina una tarea por su ID")
	fmt.Println("  mark-in-progress <id>          - Marca una tarea como 'in-progress'")
	fmt.Println("  mark-done <id>                 - Marca una tarea como 'done'")
}

// AddTask agrega una nueva tarea a la lista y retorna la lista actualizada.
func AddTask(tasksList []Task, description string) []Task {

	// Generar el ID automáticamente
	var newID int
	if len(tasksList) > 0 {
		newID = tasksList[len(tasksList)-1].ID + 1
	} else {
		newID = 1
	}

	// Crear una nueva tarea
	newTask := model.Task{
		ID:          newID,
		Description: description,
		Status:      "todo", // Se asume que es un estado inicial
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// Añadir la nueva tarea a la lista
	return append(tasksList, newTask)
}

func ListTask(taskList []Task) {
	if len(taskList) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range taskList {
		fmt.Printf("Id: %d | Description: %s | Status: %s | CreatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func ListTaskDone(taskList []Task, status string) {
	if len(taskList) == 0 {
		fmt.Println("No hay tareas")
		return
	}
	for _, task := range taskList {
		if task.Status == status {
			fmt.Printf("Id: %d | Description: %s | Status: %s | CreatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"))

		}
	}
}

func ListTaskTodo(taskList []Task, status string) {
	if len(taskList) == 0 {
		fmt.Println("No hay tareas")
		return
	}
	for _, task := range taskList {
		if task.Status == status {
			fmt.Printf("Id: %d | Description: %s | Status: %s | CreatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"))

		}
	}
}

func ListInProgress(taskList []Task, status string) {
	if len(taskList) == 0 {
		fmt.Println("No hay tareas")
		return
	}
	for _, task := range taskList {
		if task.Status == status {
			fmt.Printf("Id: %d | Description: %s | Status: %s | CreatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"))

		}
	}
}

func UpdateTask(taskList []Task, id int, description string) bool {
	for i, task := range taskList {
		if task.ID == id {
			taskList[i].Description = description
			taskList[i].UpdatedAt = time.Now()
			return true
		}

	}
	fmt.Println("Tarea no encontrada con el ID especificado.")
	return false
}

func DeleteTask(taskList *[]Task, id int) bool {
	for i, task := range *taskList {
		if task.ID == id {
			*taskList = append((*taskList)[:i], (*taskList)[i+1:]...)
			return true // Indica que la tarea fue eliminada
		}
	}
	return false // Indica que la tarea no fue encontrada
}

func MarkTaskInProgress(taskList []Task, id int) bool {
	for i, task := range taskList {
		if task.ID == id {
			taskList[i].Status = "in-progress"
			taskList[i].UpdatedAt = time.Now()
			return true
		}

	}
	fmt.Println("Tarea no encontrada con el ID especificado.")
	return false
}

func MarkTaskAsDone(taskList []Task, id int) bool {
	for i, task := range taskList {
		if task.ID == id {
			taskList[i].Status = "done"
			taskList[i].UpdatedAt = time.Now()
			return true
		}

	}
	fmt.Println("Tarea no encontrada con el ID especificado.")
	return false
}

// Funcion para escribir en el archivo .json
func SaveTask(fileName string, tasksList []Task) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.MarshalIndent(tasksList, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	return err

}
