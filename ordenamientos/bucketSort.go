package ordenamientos

// Podemos ordenar algo que puede no tener un rango enumerable (discreto)
// Debe ser conocida la distribución de los datos
// datos uniformemente distribuidos

// Estabilidad: depende del ordenamiento interno
// In-place: ya vemos que no, por separar en baldes
import (
	"sort"
)

func BucketSort(arr []float64) []float64 {
	if len(arr) <= 1 {
		return arr
	}
	n := len(arr)

	// 1. Crear n buckets (listas vacías)
	buckets := make([][]float64, n)

	// 2. Colocar cada elemento en el bucket correspondiente
	for _, val := range arr {
		indice := int(val * float64(n))
		if indice == n {
			indice = n - 1
		}
		buckets[indice] = append(buckets[indice], val)
	}
	for i := range buckets {
		sort.Float64s(buckets[i])
	}
	res := make([]float64, 0)
	for _, bucket := range buckets {
		res = append(res, bucket...)
	}
	return res
}

// Crear Baldes: O(b)
// Separar en cada balde: O(n)
// Pasar por cada Balde, ordenando: si está bien distribuido, podemos dejarlo en O(b + n)
// Juntar: O(b + n)

// Complejidad Total: O(b + n)

func BucketSortSimpleCifras(arr []int, cifra int) []int {
	if len(arr) <= 1 {
		return arr
	}
	max := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	cantBuckets := max
	buckets := make([][]int, cantBuckets)
	for _, n := range arr {
		digito := (n / cifra) % 10
		buckets[digito] = append(buckets[digito], n)
	}
	result := make([]int, 0)
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}
	return result

}
