package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades recursivo
// Devuelve un slice con las actividades seleccionadas que no se solapan
// Pre condición: las actividades están ordenadas  de menor a mayor por tiempo de finalización
// Reescribir la función que resuelve el problema de seleccionar actividades en forma recursiva
func SelectorRecursivo(actividades []Actividad) []Actividad {
	seleccionadas := []Actividad{}

	if len(actividades) == 0 {
		return []Actividad{}
	}

	//Seleccionamos la actividad de menor fin, es la primera porque ya esta ordenado.
	seleccionadas = append(seleccionadas, actividades[0])

	// Nos quedamos con las actividades que no se solapan con la actual.
	aux := []Actividad{}
	for i := 1; i < len(actividades); i++ {
		if actividades[i].Inicio >= actividades[0].Fin {
			aux = append(aux, actividades[i])
		}
	}
	// Llamada recursiva
	seleccionadas = append(seleccionadas, SelectorRecursivo(aux)...)

	return seleccionadas
}
