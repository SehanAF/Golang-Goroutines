package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Ini adalah Goroutine
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

// Kode : Membuat banyak Goroutine
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i) // Dengan menambahkan kata kunci "go"
	}

	time.Sleep(400 * time.Second) // Menambahkan time.sleep adalah optional.
}
