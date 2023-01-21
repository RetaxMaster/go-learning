package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	// Reemplaza todos los elementos por todos los elementos antes del indice y todos los elementos después del indice (aisla al que se elimina y simplemente no lo agrega)
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleto() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {

	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de Platzi en esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de Platzi en esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de NodeJS",
		descripcion: "Completar mi curso de NodeJS de Platzi en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarALista(t3)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i, "Nombre", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks {
		fmt.Println("Index", index, "Nombre", tarea.nombre)
	}

}
