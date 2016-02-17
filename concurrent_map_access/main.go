// Run this with a go 1.5 compiler and you will either get an error
// message that may not be helpful or consistent ('overflow is not
// nil', 'bad map state', 'nil pointer dereference', etc.). With a 1.6
// compiler you will get a 'concurrent map read and map write' error
// message. You can then use the race detector to get more info on the
// problem.
package main

import (
	"fmt"
	"sync"
)

func main() {
	// reduce this number and see the effect
	const workers = 99

	var wg sync.WaitGroup
	wg.Add(workers)
	m := map[int]int{}
	// we are going to spin up a bunch of go routines that write
	// to the same map a bunch
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	// read from a map that is also being written to by a ton of
	// go routines
	fmt.Println(m)
}
