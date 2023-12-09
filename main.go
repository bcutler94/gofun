package main

import (
	"fmt"
	"log"
)

func main() {
	keys := make(chan string)

	for i := 0; i < 5; i++ {
		go func (i int)  {
			keys <- fmt.Sprintf("%d", i)
		}(i)
	}

	for i := 0; i < 5; i++ {
		log.Println(<-keys)
	}

}
