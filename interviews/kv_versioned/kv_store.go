package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

// version/value pair
type versionedValue struct {
	version int64
	value   int
}

// KeyValueStore is a KeyValue store with global versioning
type KeyValueStore struct {
	version int64
	values  map[string][]versionedValue // mapped by key
	//lock    sync.Mutex // not used for now, we could use if we want to allow parallel operation on the map
}

// NewKeyValueStore creates a new KeyValue store
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		version: 0,
		values:  map[string][]versionedValue{},
	}
}

// Put puts a value for a key
func (kv *KeyValueStore) Put(key string, value int) {

	version := atomic.AddInt64(&kv.version, 1)

	if _, found := kv.values[key]; !found {
		kv.values[key] = []versionedValue{}
	}
	kv.values[key] = append(kv.values[key], versionedValue{version, value})

	fmt.Printf("PUT(#%d) %s %d\n", version, key, value)
}

// Get returns the latest value of a key
func (kv *KeyValueStore) Get(key string) {
	values, found := kv.values[key]
	if !found {
		fmt.Printf("GET %s = <NULL>\n", key)
		return
	}
	latestValue := values[len(values)-1].value
	fmt.Printf("GET %s = %d\n", key, latestValue)
}

// GetVersioned returns the vakue of a key for a given version
func (kv *KeyValueStore) GetVersioned(key string, version int) {
	values, found := kv.values[key]
	if !found {
		fmt.Printf("GET %s(#%d) = <NULL>\n", key, version)
		return
	}
	var foundVal *int
	for _, v := range values {
		if v.version > int64(version) {
			break // don't consider newer values
		}
		foundVal = &v.value
	}
	if foundVal == nil {
		fmt.Printf("GET %s(#%d) = NULL\n", key, version)
		return
	}
	fmt.Printf("GET %s(#%d) = %d\n", key, version, *foundVal)
}

// Process processes an input file
func (kv *KeyValueStore) ProcessInput(inputFile string) error {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Invalid file %s", inputFile)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PUT") {
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return fmt.Errorf("Invalid input %s", line)
			}
			value, err := strconv.Atoi(parts[2])
			if err != nil {
				return fmt.Errorf("Invalid input %s", line)
			}
			kv.Put(parts[1], value)
		} else if strings.HasPrefix(line, "GET") {
			parts := strings.Fields(line)
			if len(parts) != 2 && len(parts) != 3 {
				return fmt.Errorf("Invalid input %s", line)
			}
			if len(parts) == 2 {
				kv.Get(parts[1])
			} else {
				version, err := strconv.Atoi(parts[2])
				if err != nil {
					return fmt.Errorf("Invalid input %s", line)
				}
				kv.GetVersioned(parts[1], version)
			}
		}
	}
	return nil
}
