package ordenamientos

// Elementos con varios dígitos o componentes, con misma cantidad de cifras, o muy similar..
// Ordenamiento auxiliar que TIENE que ser estable si no no funciona
// Ordena de la cifra menos significativa a la más significativa.

func RadixSort(arr []int, aux string) []int {
	//Ordenamos primero por cifra menos significativa
	//Ordenamos por la siguiente cifra... hasta llegr a la ultima cifra, la que queremos que quede
	// ordenada
	max := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	if aux == "bucket" {
		for i := 1; max/i > 0; i *= 10 {
			arr = BucketSortSimpleCifras(arr, i)
		}
	}
	return arr
}
