package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// dua cara deklarasi mutex

// mutex berlabel
var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)

// anonymous mutex
// var cond = sync.NewCond(&sync.mutex{})

var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	fmt.Println("Pemenang lap 1", value)
	cond.Wait()
	fmt.Println("Pemenang lap 2", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
	fmt.Println("uji coba selesai")
}

/*
	perbedaan utama antara cond dan locking biasa adalah METHOD WAIT
	setelah goroutines berhasil melakukan locking terhadap sesuatu, apa yang terjadi selanutnya
	bukanlah eksekusi, tetapi wait condition. Aritnya jangan langunsg lanjut, tetapi tunggu dulu
*/
