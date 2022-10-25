package integers

import (
	"testing"
	"fmt"
)

// Funcao que trata os erros e retorna uma string
func TratamentoDeCondicao(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot = %d\nwant = %d", got, want)
	}
}

// Teste da funcao de soma
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	TratamentoDeCondicao(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}