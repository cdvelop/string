package strings_test

import (
	"testing"

	"github.com/cdvelop/strings"
)

var stringTest = strings.Strings{}

func TestToUpperCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"4 hello! 123", "4 HELLO! 123"},
		{"world 52,", "WORLD 52,"},
		{"123", "123"},
		{"", ""},
		{"hola mundo", "HOLA MUNDO"},
		{"hola ñurdo", "HOLA ÑURDO"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := stringTest.ToUpperCase(test.input)

			if result != test.expected {
				t.Fatalf("Entrada: %s\n-Respuesta: %s\n-Expectativa: %s", test.input, result, test.expected)
				return
			}
		})
	}

}

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"4 HELLO! 123", "4 hello! 123"},
		{"World 52,", "world 52,"},
		{"123", "123"},
		{"", ""},
		{"HOLA MUNDO", "hola mundo"},
		{"HOLA Ñurdo", "hola ñurdo"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := stringTest.ToLowerCase(test.input)

			if result != test.expected {
				t.Fatalf("Entrada: %s\n-Respuesta: %s\n-Expectativa: %s", test.input, result, test.expected)
				return
			}
		})
	}
}

func TestToUpperCaseFirsLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hELLO", "HELLO"},
		{"world 52,", "World 52,"},
		{"123", "123"},
		{"", ""},
		{"hOLA MUNDO", "HOLA MUNDO"},
		{"ñurdo hello", "Ñurdo hello"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := stringTest.UpperCaseFirstLetter(test.input)

			if result != test.expected {
				t.Fatalf("Entrada: %s\n-Respuesta: %s\n-Expectativa: %s", test.input, result, test.expected)
				return
			}
		})
	}
}

func TestLowerCaseFirstLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO ", "hELLO "},
		{"world 52,", "world 52,"},
		{"123", "123"},
		{"", ""},
		{"HOLA MUNDO", "hOLA MUNDO"},
		{"Ñurdo hello", "ñurdo hello"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := stringTest.LowerCaseFirstLetter(test.input)

			if result != test.expected {
				t.Fatalf("Entrada: %s\n-Respuesta: %s\n-Expectativa: %s", test.input, result, test.expected)
				return
			}
		})
	}
}
