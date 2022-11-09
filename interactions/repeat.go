package interactions

func Repeat(char string, qtd int) string {
	var repeated string

	for i := 0; i < qtd; i++ {
		repeated += char
	}

	return repeated
}