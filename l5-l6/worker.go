package main

import (
	"fmt"
	"strings"
)

func Map(key string, value string) []KVPair {

	output := make([]KVPair, 0)

	// TODO: This loop iterates over each line of the "value" string
	// You will want to parse out the date and temperature from each line and add it to the "output" slice
	for _, line := range strings.Split(strings.TrimSuffix(value, "\n"), "\n") {
		fmt.Println(line)
		fields := strings.Split(line, ",")
		year := fields[1]
		year = year[0:4]
		temp := fields[2]
		output = append(output, KVPair{year, temp})

	}

	return output
}

//func Reduce(key string, value []string) float64 {

// Converting from a string to float may be useful
//val, err := strconv.ParseFloat(INPUT, 64)

//}
