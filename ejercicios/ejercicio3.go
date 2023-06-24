package ejercicios

import (
	"sort"
)

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

// Problema de la mochila 0/1
// Se tienen n objetos (cada objeto i tiene un peso y un valor);
// y una mochila con capacidad máxima de W.
// Se pretende encontrar la manera de cargar la mochila de forma que se maximice el valor de lo transportado
// y se respete su capacidad máxima. Los objetos se pueden fraccionar de tal forma de cargar solo una fracción del objeto

// > El paso greedy es elegir primero el elemento que tenga mayor valor por unidad de peso

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	aux := make(map[Objeto]float32)
	capacidad2 := 0
	objetos = ordenarObjetos(objetos)
	for i := 0; i < len(objetos); i++ {
		if (objetos[i].Valor / objetos[i].Peso) >= (objetos[i+1].Valor / objetos[i+1].Peso) {
			aux[objetos[i]] = float32(objetos[i].Peso)
			capacidad2 += objetos[i].Peso
		}

	}
	return aux

}

func ordenarObjetos(objetos []Objeto) []Objeto {

	sort.Slice(objetos, func(i, j int) bool {
		return objetos[i].Valor > objetos[j].Valor
	})

	return objetos
}
