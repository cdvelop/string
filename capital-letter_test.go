package strings_test

import "testing"

func TestFirstLetterIsCapital(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"4 hello! 123", false},
		{"world 52,", false},
		{"World 52,", true},
		{"123", false},
		{"", false},
		{"HOLA MUNDO", true},
		{"HOLA Ñurdo", true},
		{"Ñandú", true},
		{" Ñandú con espacio al inicio", true},
		{"Ál inicio con tilde", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := stringTest.FirstLetterIsCapital(test.input)

			if result != test.expected {
				t.Fatalf("\nEntrada: %s\n-Respuesta: %v\n-Expectativa: %v", test.input, result, test.expected)
				return
			}
		})
	}

}
