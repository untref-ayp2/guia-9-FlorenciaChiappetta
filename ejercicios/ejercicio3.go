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
	objetovalpes := []int{}
	var peso int
	var valor int
	var pesorestante int 

	mapa := make(map[Objeto]float32)
	for i:= 0; i < len(objetos); i++{
	objetovalpes[i] = objetos[i].Valor / objetos[i].Peso
	}

	//ordenamos el arreglo por burbujeo de mayor a menor
	for i := 0; i < len(objetos);i++ {
		for j:=i ; j < len(objetos)-1-i;i++{
			if objetovalpes[j] > objetovalpes[j+1] {
				 objetovalpes[j],objetovalpes[j+1] = objetovalpes[j+1],objetovalpes[j]
			}  
		}

		if objetos[i].Peso <= capacidad {			
			peso += objetos[i].Peso
			valor += objetos[i].Valor
		}
		}
	}

} 