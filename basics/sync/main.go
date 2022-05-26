package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	sync.Mutex
	m map[string]string
}

func (m *SafeMap) Store(key, value string) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func (m *SafeMap) Load(key string) (string, bool) {
	m.Lock()
	defer m.Unlock()
	val, ok := m.m[key]
	return val, ok
}

func (m *SafeMap) Delete(key string) {
	m.Lock()
	delete(m.m, key)
	m.Unlock()
}

func main() {
	m := SafeMap{
		m: map[string]string{},
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Store("test", "Hello world 1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Store("test", "Hello world 2")
	}()

	wg.Wait()

	fmt.Println(m.Load("test"))
	m.Delete("test")
	fmt.Println(m.Load("test"))
}
