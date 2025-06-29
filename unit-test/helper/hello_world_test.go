package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
