package main

import (
	"./lrucache"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type getter func(key interface{}) (interface{}, bool)

func printData(get getter, key interface{}) {
	if val, exists := get(key); !exists {
		fmt.Println("NULL")
	} else {
		fmt.Println(val)
	}
}

func dumper(key, val interface{}) {
	fmt.Println(key, val)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	lru := lrucache.New()

	for i := 0; i < N && scanner.Scan(); i++ {
		words := strings.Fields(scanner.Text())
		// fmt.Println("get words:", words)

		switch words[0] {
		case "BOUND":
			nb, _ := strconv.Atoi(words[1])
			lru.SetBound(nb)
		case "SET":
			lru.Set(words[1], words[2])
		case "GET":
			printData(lru.Get, words[1])
		case "PEEK":
			printData(lru.Peek, words[1])
		case "DUMP":
			lru.Do(dumper)
		}
	}
}
