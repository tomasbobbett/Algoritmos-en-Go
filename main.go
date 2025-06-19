package main

import (
	o "algoritmosengo/ordenamientos"
	"fmt"
)

// Un algoritmo no comparativo no puede ser mejor que O(n log(n))
func main() {
	arr := []int{4, 2, 2, 8, -1, 3, 3, 1, -2, -4, -8}
	fmt.Println("Arreglo original:", arr)

	fmt.Println("Arreglo ordenado:", o.Counting(arr))
}
