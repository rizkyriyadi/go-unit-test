package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yamdi")

	if result != "Yamdi, Selamat datang" {
		// error
		t.Error("Harusnya Yamdi, Selamat datang")
	}
}

func TestHelloWorldAsstert(t *testing.T) {
	result := HelloWorld("Gaji")

	assert.Equal(t, "Gaji, Selamat datangg", result)
	fmt.Println("End of TestHelloWorldAssert")
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE TEST BOSKU")

	m.Run()

	fmt.Println("AFTER TEST BOSKU")

}
