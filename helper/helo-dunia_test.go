package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yamdi")

	if result != "Yamdi, Selamat datang" {
		// error
		t.Error("Harusnya Yamdi, Selamat datang")
	}
}
