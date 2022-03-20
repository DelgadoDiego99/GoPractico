package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tk *task) {
	t.tasks = append(t.tasks, tk)
}

func (t *taskList) eliminarTarea(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, valor := range t.tasks {
		fmt.Println("Nombre:", valor.nombre)
		fmt.Println("Descripción:", valor.descripcion)
		fmt.Println("Completado:", valor.completado)
	}
}

func (t *taskList) imprimirCompletadas() {
	for _, valor := range t.tasks {
		if valor.completado == true {
			fmt.Println("Nombre:", valor.nombre)
			fmt.Println("Descripción:", valor.descripcion)
			fmt.Println("Completado:", valor.completado)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompletado() {
	t.completado = true
}

func (t *task) actualizarDescripcion(desc string) {
	t.descripcion = desc
}

func (t *task) actualizarNombre(name string) {
	t.nombre = name
}

func main() {

	mapaTareas := make(map[string]*taskList)

	t1 := &task{
		nombre:      "Completar curso",
		descripcion: "Completar el curso entre hoy y mañana",
		completado:  false,
	}
	t2 := &task{
		nombre:      "Completar curso de Py",
		descripcion: "Completar el curso este semestre",
		completado:  false,
	}
	t3 := &task{
		nombre:      "Completar curso de Go",
		descripcion: "Completar el curso mañana",
		completado:  false,
	}

	listaTareas := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	listaTareas.agregarALista(t3)
	listaTareas.tasks[0].marcarCompletado()
	/*fmt.Println("Tareas completadas")
	listaTareas.imprimirCompletadas()*/

	mapaTareas["Prioridad Alta"] = listaTareas

	t4 := &task{
		nombre:      "Sacar la basura",
		descripcion: "Toca sacar la basura a mas tardar el mediodia",
		completado:  false,
	}
	t5 := &task{
		nombre:      "Leer un libro",
		descripcion: "Toca leer 25 páginas",
		completado:  false,
	}

	listaTareasCasa := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Prioridad Media"] = listaTareasCasa

	fmt.Println("Tareas de cursos")
	mapaTareas["Prioridad Alta"].imprimirLista()

	fmt.Println("Tareas de casa")
	mapaTareas["Prioridad Media"].imprimirLista()

}
