package belajar_golang_context

import (
	"fmt"
	"testing"
	"time"
)

func channelwithoutdata() chan int {
	channel1 := make(chan int)
	go func() {
		defer close(channel1)
		time.Sleep(5 * time.Second)
		makeDataForChannel(channel1)
	}()

	return channel1
}

func makeDataForChannel(channel chan int) {
	number := 1312
	channel <- number
}

func TestChannelWithoutData(t *testing.T) {
	destination := channelwithoutdata()
	fmt.Println("menunggu data channel dikirimkan")
	fmt.Println(<-destination)
}

// ternyata benar bahwa goroutines yang membaca channel kosong akan menunggu hingga ada data yang dikirm
