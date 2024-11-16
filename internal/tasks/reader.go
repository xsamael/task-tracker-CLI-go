package tasks

import (
	"encoding/json"
	"io"
	"os"
)

// ReadJson lee el archivo JSON y lo deserializa en la lista de tareas.
func ReadJson(filename string, taskList *[]Task) error {
	// Abrir o crear archivo
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Verificar si el archivo está vacío
	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.Size() > 0 {
		// Leer contenido del archivo
		bytes, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		// Deserializar JSON
		err = json.Unmarshal(bytes, taskList)
		if err != nil {
			return err
		}
	} else {
		*taskList = []Task{}
	}

	return nil
}
