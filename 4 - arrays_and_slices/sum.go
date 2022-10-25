package main

func Sum(numeros []int) int {
	soma := 0
	for _, elemento := range numeros {
		soma += elemento
	}
	return soma
}