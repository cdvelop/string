package strings_test

import (
	"testing"

	"github.com/cdvelop/strings"
)

func TestToLowerCaseAlphabet(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO", "hello"},
		{"World 52,", "world 52,"},
		{"123", "123"},
		{"", ""},
		{"HOLA MUNDO", "hola mundo"},
		{"HOLA Ñurdo", "hola ñurdo"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := strings.ToLowerCase(test.input)

			if result != test.expected {
				t.Fatalf("Entrada: %s\n-Respuesta: %s\n-Expectativa: %s", test.input, result, test.expected)
				return
			}
		})
	}
}
