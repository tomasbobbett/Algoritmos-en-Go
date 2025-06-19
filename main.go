package main

import (
	o "algoritmosengo/ordenamientos"
	"fmt"
)

// Un algoritmo no comparativo no puede ser mejor que O(n log(n))
func main() {
	arr := []int{4, 2, 2, 8, -1, 3, 3, 1, -2, -4, -8}
	fmt.Println("Arreglo original:", arr)
	fmt.Println("Counting Sort:", o.Counting(arr))

	arr2 := []int{703, 142, 112, 712, 507, 555, 505, 205}
	fmt.Println("\n\nArreglo original:", arr2)
	fmt.Println("Radix Sort:", o.RadixSort(arr2, "bucket"))

	arr3 := []float64{0.1, 0.02, 0.5, 0.54, 0.95, 0.02, 0.78}
	fmt.Println("\n\nArreglo original:", arr3)
	fmt.Println("Bucket Sort:", o.BucketSort(arr3))
}
