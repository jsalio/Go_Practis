package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	TaskName string
	Status   string
	Date     string
}

func loadTasksFromFile(tasks *[]Task) error {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if erred := decoder.Decode(tasks); erred != nil {
		return fmt.Errorf("error al leer el archivo: %v", erred)
	}
	return nil
}

func saveTasksToFile(tasks *[]Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return fmt.Errorf("error al crear el archivo: %v", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}
	return nil
}

func findTaskIndex(slice []Task, pred func(string) bool) (int, bool) {
	found := false
	index := -1
	for sliceIndex, item := range slice {
		if pred(item.TaskName) {
			index = sliceIndex
			found = true
		}
	}
	return index, found
}

func AddTask(tasks *[]Task) {
	name := ""
	println("Agregar una nueva tarea :")
	fmt.Print("Nombre de la tarea :")
	fmt.Scanln(&name)
	var newTask Task
	newTask.TaskName = name
	newTask.Date = ""
	newTask.Status = "Created"
	*tasks = append(*tasks, newTask)
}

func markTaskAsComplete(tasks *[]Task) {
	if len(*tasks) == 0 {
		print("La lista de tareas esta vacia")
		return
	}
	println("Agregar una nueva tarea :")
	finder := ""
	print("Ingrese el nombre de la tarea :")
	fmt.Scanln(&finder)
	index, found := findTaskIndex(*tasks, func(s string) bool {
		return s == finder
	})
	if !found {
		println("Task not found")
		return
	}
	(*tasks)[index].Status = "complete" //asi es como se actuliza una referencia
}

func RemoveTask(tasks *[]Task) {
	if len(*tasks) == 0 {
		print("La lista de tareas esta vacia")
		return
	}
	println("Eliminar tarea :")
	finder := ""
	print("Ingrese el nombre de la tarea :")
	fmt.Scanln(&finder)
	index, found := findTaskIndex(*tasks, func(s string) bool {
		return s == finder
	})
	if !found {
		println("Task not found")
		return
	}
	(*tasks)[index] = (*tasks)[len(*tasks)-1] // Reemplaza con el último elemento
	*tasks = (*tasks)[:len(*tasks)-1]         // Reduce el slice

}

func print_tasks(tasks *[]Task) {
	println("Listando tareas")
	for _, task := range *tasks {
		fmt.Printf("Tarea : {%s} estado [%s] \n", task.TaskName, task.Status)
	}
}

func app_menu() string {
	option := ""
	println("Menu de aplicacion :")
	println("[1] agregar")
	println("[2] completar")
	println("[3] listar")
	println("[4] eliminar tarea")
	println("[5] salir")
	println()
	println("Elija una opcion (1-4) :")
	fmt.Scanln(&option)
	return option
}

func main() {
	tasks := []Task{} //Debo repasar esto con los slice
	if err := loadTasksFromFile(&tasks); err != nil {
		fmt.Println("Error al cargar las tareas:", err)
	}
	for {
		option := app_menu()
		switch option {
		case "1":
			AddTask(&tasks)
		case "2":
			markTaskAsComplete(&tasks)
		case "3":
			print_tasks(&tasks)
		case "5":
			if err := saveTasksToFile(&tasks); err != nil {
				fmt.Println("Error al guardar las tareas:", err)
			} else {
				println("Tareas guardadas exitosamente")
			}
			println("Saliendo de la aplicación")
			return
		case "4":
			RemoveTask(&tasks)
		default:
			println("opcion invalida")
		}
	}
}
