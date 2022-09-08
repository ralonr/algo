package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]string{"key1": "val1", "key2": "val2", "key3": "val3", "key4": "val4"}
	sortStringMap(m)
}

// sortStringMap prints the [string]string as keys sorted
func sortStringMap(m map[string]string) {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys) // sort the keys
	for _, key := range keys {
		fmt.Printf("%s\t:%s\n", key, m[key])
	}
}
