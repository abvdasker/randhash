package main

import (
	"fmt"
	"github.com/abvdasker/randhash/lib"
)

func main() {
	hash := lib.NewRandHash()
	hash.Put("test", "first")
	hash.Put("test2", "second")
	hash.Put("test3", "third")
	printRandom(hash)
}

func printString(hash *lib.RandHash, key interface{}) {
	raw := hash.Get(key)
	value, _ := raw.(string)
	fmt.Println(value)
}

func printRandom(hash *lib.RandHash) {
	raw := hash.Sample()
	value := raw.(string)
	fmt.Println(value)
}
