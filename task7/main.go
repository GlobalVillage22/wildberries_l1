package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

// Создание структуры MyMap с мютексом
type MyMap struct {
	data map[int]int
	mx   sync.Mutex
}

func NewMap() *MyMap {
	return &MyMap{data: make(map[int]int)}
}

// Метод Set
func (m *MyMap) Set(key, val int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.data[key] = val
}

// Метод Get
func (m *MyMap) Get(key int) (int, bool) {
	m.mx.Lock()
	defer m.mx.Unlock()
	val, ok := m.data[key]
	return val, ok
}
func (m *MyMap) RangePrint() {
	m.mx.Lock()
	defer m.mx.Unlock()
	for i, v := range m.data {
		fmt.Printf("key: %v, val: %v \n", i, v)
	}
}

func main() {
	mymap := NewMap()
	wg := sync.WaitGroup{}
	msync := sync.Map{} // Map из библиотеки sync
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			mymap.Set(i, rand.Int())
			msync.Store(i, rand.Int())
		}()
	}
	wg.Wait()
	mymap.RangePrint()
	fmt.Println("----------------")
	msync.Range(func(key, value any) bool {
		fmt.Printf("key: %v, val: %v \n", key, value)
		return true
	})

}
