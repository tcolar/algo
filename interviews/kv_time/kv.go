package main

import (
	"errors"
	"fmt"
	"time"
)

var notFoundError = errors.New("Key not found")

type Kv struct {
	Data map[string][]ValueAtTime
}

type ValueAtTime struct {
	Time  int64
	Value string
}

func NewKv() *Kv {
	return &Kv{
		Data: map[string][]ValueAtTime{},
	}
}

func (kv *Kv) Set(at int64, key, value string) {
	kv.Data[key] = append(kv.Data[key], ValueAtTime{
		Time:  at,
		Value: value,
	})
}

func (kv *Kv) Get(key string) (string, error) {
	if array, found := kv.Data[key]; found {
		return array[len(array)-1].Value, nil
	}
	return "", fmt.Errorf("Key not found %s", key)
}

func (kv *Kv) GetAt(key string, atTime int64) (string, error) {
	if array, found := kv.Data[key]; found {
		var value string
		var valueFound bool
		for _, valueAtTime := range array {
			if valueAtTime.Time > atTime {
				break
			}
			value = valueAtTime.Value
			valueFound = true
		}
		if valueFound {
			return value, nil
		}
	}
	return "", fmt.Errorf("Key not found %s at %d", key, atTime)
}

func now() int64 {
	return time.Now().UnixNano()
}

func main() {
	kv := NewKv()
	kv.Set(1, "name", "John")
	kv.Set(5, "name", "Bob")
	kv.Set(7, "age", "24")
	fmt.Println(kv.Get("name"))
	fmt.Println(kv.Get("bar"))
	fmt.Println(kv.GetAt("name", 0))
	fmt.Println(kv.GetAt("name", 1))
	fmt.Println(kv.GetAt("name", 4))
	fmt.Println(kv.GetAt("name", 5))
	fmt.Println(kv.GetAt("foo", 4))
	//fmt.Println(kv.Get("foo"))
}
