package divYConquista

// BuscarMinimo devuelve el valor del minimo del arreglo, no su posicion
// Precondicion: el arreglo tiene al menos un elemento
func BuscarMinimo(arr []int) int {

	return buscarRec(arr, 0, len(arr)-1)
}

func buscarRec(arr []int, inicio, final int) int {
	if inicio == final {
		return arr[inicio]
	}
	mid := (inicio + final) / 2
	aIzq := buscarRec(arr, inicio, mid)  // 1 llamada
	aDer := buscarRec(arr, mid+1, final) //  2 llamadas A = 2
	if aIzq < aDer {
		return aIzq
	}
	return aDer
}

// a = 2
// b = 2			log2(2) = 1 > c => O(n)
// c = 0
