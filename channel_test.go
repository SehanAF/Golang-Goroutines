package golang_goroutines

import (
	"fmt"
	"strconv"
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

func OnlyIn(channel chan<- string) {
	time.Sleep(3 * time.Second)
	channel <- "Sehan Alfarisi"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Ini adalah buffered menggunakan channel.
/**func TestBuffredChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Sehan"
	channel <- "Alfarisi"
	channel <- "Gavi"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}
*/
// Ini adalah buffered goroutine.
func TestBuffredChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Sehan"
		channel <- "Alfarisi"
		channel <- "Gavi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")

}

// Ini adalah Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// Ini dengan cara ribet dengan harus membuat 2 select.
	/**select {
	case data := <-channel1:
		fmt.Println("Data dari channel 1", data)
	case data := <-channel2:
		fmt.Println("Data dari channel 2", data)
	}

	select {
	case data := <-channel1:
		fmt.Println("Data dari channel 1", data)
	case data := <-channel2:
		fmt.Println("Data dari channel 2", data)
	}

	// ini dengan cara simple dengan menggunakan perulangan
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}*/

	// ini dengan menggunakan select
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}
