// Running this with a go 1.5 compiler will yield about 600 lines of
// stderr output because it prints the stack trace of every go routine
// running. Running this with a go 1.6 compiler will only show the
// ouput of the go routine where the panic occurred
package main

import (
	"sync"
)

func main() {
	// want to start up a bunch of go routines so there a bunch
	// running when we trigger a panic
	var ready sync.WaitGroup
	const workers = 99

	ready.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			ready.Done()
			select {}
		}()
	}
	ready.Wait()
	var i *int
	// i is nil here, so this will trigger nil pointer dereference
	// error
	*i++
}
