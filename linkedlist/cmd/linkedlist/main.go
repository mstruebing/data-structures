package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mstruebing/data-structures/linkedlist"
)

func insert(playlist *linkedlist.Node, item string) {
	start := time.Now()

	playlist.Insert(item)

	fmt.Println(fmt.Sprintf("%s: %v", item, time.Since(start)))
}

func main() {
	playlist := linkedlist.CreateEmpty()

	for i := 0; i < 1000; i++ {
		insert(playlist, strconv.Itoa(i))
	}

	// playlist.Print()
}
