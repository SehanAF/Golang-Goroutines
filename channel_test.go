package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Ini adalah kode: Channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // Ini adalah kode: Membuat Channel.
	defer close(channel)

	go func() {
		time.Sleep(5 * time.Second)
		channel <- "Muhammad Sechan Alfarisi"
		fmt.Println("Selesai mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(10 * time.Second)
}

// Dibawah ini adalah Kode: Channe; Sebagai Parameter.
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Sechan Alfarisi"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(10 * time.Second)
}
