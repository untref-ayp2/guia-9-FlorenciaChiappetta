package actividades

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades iterativo
// Devuelve un slice con las actividades seleccionadas que no se solapan
// enfoque greedy
// Pre condición: las actividades están ordenadas  de menor a mayor por tiempo de finalización
func SelectorActividadesIterativo(actividades []Actividad) []Actividad {
	resultado := []Actividad{}
	if len(actividades) == 0 {
		return []Actividad{}
	}

	resultado = append(resultado, actividades[0])
	aux := []Actividad{}
	for i := 0; i < len(actividades); i++ {
		for j := i; j < len(actividades)-i-1; i++ {
			if actividades[i].Inicio > actividades[0].Fin {
				aux = append(aux, actividades[i])
			}
		}
	}
	resultado = append(resultado, SelectorActividadesIterativov1(aux)...)
	return resultado
}
func SelectorActividadesIterativov1(actividades []Actividad) []Actividad {
	var seleccionadas []Actividad
	n := len(actividades)
	seleccionadas = append(seleccionadas, actividades[0])
	k := 0
	for i := 1; i < n; i++ {
		if actividades[i].Inicio >= actividades[k].Fin {
			seleccionadas = append(seleccionadas, actividades[i])
			k = i
		}
	}
	return seleccionadas
}
