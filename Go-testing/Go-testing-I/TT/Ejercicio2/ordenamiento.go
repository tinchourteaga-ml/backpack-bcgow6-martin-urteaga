package Ejercicio2

import "sort"

func OrdenarSliceAsc(valores []int) []int {
	sort.Slice(valores, func(i, j int) bool {
		return valores[i] < valores[j]
	})
	return valores
}
