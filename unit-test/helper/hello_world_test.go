package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) { // suitable for testing responses on different variable types
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Noby",       // test name
			request:  "Noby",       // input to the function
			expected: "Hello Noby", // expected output
		},
		{
			name:     "Bitzer",
			request:  "Bitzer",
			expected: "Hello Bitzer",
		},
		{
			name:     "Adata",
			request:  "Adata",
			expected: "Hello Adata",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Noby", func(t *testing.T) {
		result := Hello("Noby")
		require.Equal(t, "Hello Noby", result, "Result must be 'Hello Noby'")
	})
	t.Run("Bitzer", func(t *testing.T) {
		result := Hello("Bitzer")
		require.Equal(t, "Hello Bitzer", result, "Result must be 'Hello Bitzer'")
	})
	t.Run("Gon", func(t *testing.T) {
		result := Hello("Gon")
		require.Equal(t, "Hello Gon", result, "Result must be 'Hello Gon'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on windows")
	}

	result := Hello("Noby")
	require.Equal(t, "Hello Noby", result, "Result must be 'Hello Noby'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := Hello("Noby")
	require.Equal(t, "Hello Noby", result, "Result must be 'Hello Noby'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := Hello("Noby")
	assert.Equal(t, "Hello Noby", result, "Result must be 'Hello Noby'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldNoby(t *testing.T) {
	result := Hello("Noby")

	if result != "Hello Noby" {
		// error
		t.Error("Result must be 'Hello Noby'")
	}

	fmt.Println("TestHelloWorldNoby Done")
}

func TestHelloWorldBitzer(t *testing.T) {
	result := Hello("Bitzer")

	if result != "Hello Bitzer" {
		// error
		t.Fatal("Result must be 'Hello Bitzer'")
	}

	fmt.Println("TestHelloWorldBitzer Done")
}
