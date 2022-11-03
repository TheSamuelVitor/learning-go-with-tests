package main

import "testing"

func TestHello(t *testing.T) {
	lista := []struct {
		nome string
		idioma string
		want string
	}{
		{"Maria", "pt", "Ola, Maria"},
		{"Samuel", "en", "Hello, Samuel"},
		{"Kaio", "es", "Hola, Kaio"},
		{"Dayanne", "fr", "Salut, Dayanne"},
		{"Mohammed", "ar", "Idioma nao encontrado"},
	}

	for _, elemento := range lista {
		t.Run("teste " + elemento.nome, func(t *testing.T){
			got := Hello(elemento.nome, elemento.idioma)
			want := elemento.want
		
			if got != want {
				t.Errorf("\ngot %v\nwant %v", got, want)
			}
		})
	}

}
