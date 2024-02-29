package strings_test

import "testing"

func TestReplace(t *testing.T) {
	tests := []struct {
		input, old, newStr, expected string
	}{
		{"Este es un ejemplo de texto de prueba.", "ejemplo", "cambio", "Este es un cambio de texto de prueba."},
		{"Hola mundo!", "mundo", "Gophers", "Hola Gophers!"},
		{"abc abc abc", "abc", "123", "123 123 123"},
		{"abc", "xyz", "123", "abc"},
		{"", "", "123", ""},
		{"abcdabcdabcd", "cd", "12", "ab12ab12ab12"},
		{"palabra, punto,", ",", ".", "palabra. punto."},
	}

	for _, test := range tests {
		result := stringTest.Replace(test.input, test.old, test.newStr)
		if result != test.expected {
			t.Errorf("Para input '%s', old '%s', new '%s', esperado '%s', pero obtenido '%s'", test.input, test.old, test.newStr, test.expected, result)
		}
	}
}

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		input, suffix, expected string
	}{
		{"hello.txt", ".txt", "hello"},
		{"example", "123", "example"},
		{"file.txt.txt", ".txt", "file.txt"},
		{"", "", ""},
		{"abc", "xyz", "abc"},
	}

	for _, test := range tests {
		result := stringTest.TrimSuffix(test.input, test.suffix)
		if result != test.expected {
			t.Errorf("Para input '%s', suffix '%s', esperado '%s', pero obtenido '%s'", test.input, test.suffix, test.expected, result)
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"  hello world  ", "hello world"},
		{"abc123", "abc123"},
		{"  trim me  ", "trim me"},
		{"", ""},
		{"  ", ""},

		{"    mucho espacio\n\n\t\tcon salto\n\n\t\tde linea     \n\t\t\t\t\t\t\n\t\t\t\t", "mucho espacio\n\n\t\tcon salto\n\n\t\tde linea"},

		{`    mucho espacio
		
		con salto

		de linea     
		              
		
		`, `mucho espacio
		
		con salto

		de linea`},
	}

	for _, test := range tests {
		result := stringTest.Trim(test.input)
		if result != test.expected {
			t.Errorf("Para input '%s', esperado '%s', pero obtenido '%s'", test.input, test.expected, result)
		}
	}
}
