package ordenamientos

// Datos enumerables (discretos)
// Rango acotado y conocido (o facilmente obtenido)

// Esta implementación de Counting Sort es ESTABLE pero hay algunas que no
// Counting sort NO es in-place

func Counting(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	//Creamos arreglo de frecuencias freq
	max := arr[0]
	min := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	freq := make([]int, max-min+1)
	//1  Contamos frecuencias de numeros

	for _, n := range arr {
		freq[n-min]++
	}

	//2  Arreglo de donde comienza cada uno SUMAS ACUMULADAS con 0 en el inicio y luego en el resto de
	//posiciones voy poniendo la suma del anterior en el mismo arreglo mas esa misma posicion en freq

	starts := make([]int, max-min+1)
	for i := 1; i < len(starts); i++ {
		starts[i] = starts[i-1] + freq[i-1]
	}

	//3  Posicionar las cartas en sus lugares y sumamos uno en arreglo de comienzos cada vez q posiciono

	res := make([]int, len(arr))
	for _, n := range arr {
		res[starts[n-min]] = n
		starts[n-min]++
	}

	return res
}

//Complejidad de O(n + k) siendo o(n) la contida de frecuencias y o(k) son la cantidad de valores
// distintos que pueden tomar los elementos del arreglo
