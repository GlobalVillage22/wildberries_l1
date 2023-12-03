package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func writer(c chan int) {
	for {
		c <- rand.Intn(100)
	}
}

func reader(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	ch := make(chan int)
	go writer(ch)
	for i := 0; i < n; i++ {
		go reader(ch)
	}
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	fmt.Println("exit")
	os.Exit(1)
}
