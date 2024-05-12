package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	result := Hello("Dolly")
	if result != "Hello Dolly" {
		t.Fatal("Failed to run Hello() function")
	}
	t.Log("Success run Hello() function")
}

func TestHelloo(t *testing.T) {
	result := Hello("Dolly")
	assert.Equal(t, "Hello Dolly", result, "Result must be 'Hello Dolly'")
}

func TestHellooo(t *testing.T) {
	test := []struct {
		name    string
		request string
		expect  string
	}{
		{
			name:    "Dolly",
			request: "Dolly",
			expect:  "Hello Dolly",
		},
		{
			name:    "Indra",
			request: "Indra",
			expect:  "Hello Indra",
		},
		{
			name:    "Budi",
			request: "Budi",
			expect:  "Hello Budi",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := Hello(tt.request)
			assert.Equal(t, tt.expect, result)
		})
	}
}
