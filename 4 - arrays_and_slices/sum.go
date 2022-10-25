package main

func Sum(numeros []int) int {
	soma := 0
	for _, elemento := range numeros {
		soma += elemento
	}
	return soma
}

func SumAll(numeros ...[]int) []int {

	var sums []int

	for _, numeros := range numeros {
		sums = append(sums, Sum(numeros))
	}

	return sums
}