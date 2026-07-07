package main

import (
	"os"
	"sync"
)

type KVPair struct {
	key   string
	value string
}

func main() {
	cities := []string{"Tokyo", "Delhi", "Shanghai", "Sao_Paulo", "Mexico_City", "Cairo", "Mumbai", "Beijing", "Dhaka", "Osaka", "New_York", "Karachi", "Buenos_Aires", "Istanbul", "Kolkata", "Lagos", "Moscow", "London", "Paris", "Los_Angeles"}

	//Reads in input file for each city and calls Map function
	ch := make(chan KVPair)
	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			path := "data/" + city + ".txt"
			input, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			map_out := Map(path, string(input))
			for _, item := range map_out {
				ch <- item
			}
		}(city)
	}

	// Goroutine waits to close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// range over channel repeatedly reads from channel until it is closed
	kv_pairs := make(map[string][]string)
	for item := range ch {
		// TODO: correctly populate the map kv_pairs with the items read in on the channel "ch"
		kv_pairs[item.key] = append(kv_pairs[item.key], item.value)
	}

	// TODO: add calling Reduce tasks
	ch2 := make(chan KVPair)
	var wg2 sync.WaitGroup
	for key, value := range kv_pairs {
		wg2.Add(1)
		go func(key string) {
			defer wg2.Done()
			reduce_out := Reduce(key, value)
			for _, item := range reduce_out {
				ch2 <- item
			}

		}(key)
	}
}
