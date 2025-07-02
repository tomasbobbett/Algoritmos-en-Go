package divYConquista

func ArregloEsMagico(arr []int) bool {
	return esMagicoRec(arr, 0, len(arr)-1)
}

func esMagicoRec(arr []int, inicio, final int) bool {
	if inicio > final {
		return false
	}
	mid := (inicio + final) / 2 //Divido en dos el arreglo [b]

	if arr[mid] == mid { //todo el resto es O(1)  [C]
		return true
	} else if arr[mid] > mid {
		return esMagicoRec(arr, inicio, mid-1) //      1 llamda
	} else {
		return esMagicoRec(arr, mid+1, final) //    1 llamadas recursivas[a]
	}
}

//a = 1
//b = 2			log2(1) = 0 = c => O(log(n))
//c = 0
