package main

import "time"

/*
Реализовать собственную функцию sleep.
*/

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	sleep(5 * time.Second)
}
