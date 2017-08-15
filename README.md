### Worker-Pool

A "better" version of this [example](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/)

Simple and useful (for me) worker-pool with configurable queue size.
Usage example:

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	wp "github.com/robertogyn19/go-worker-pool"
)

type CustomJob struct {
	Number int
	wg     *sync.WaitGroup
}

func (cj *CustomJob) Start(id int) {
	work(cj.Number, id)
	cj.wg.Done()
}

var count int

func work(i, id int) {
	r := rand.Intn(100)
	d := time.Duration(r) * time.Millisecond
	fmt.Printf("(%-3d) work %-2d will sleep for %v\n", i, id, d)
	time.Sleep(d)
	count += 1
}

func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)

	jobChan := make(chan wp.GenericJob)

	dispatcher := wp.NewDispatcher(20, jobChan)
	dispatcher.Run()

	for i := 0; i < 100; i++ {
		cj := new(CustomJob)
		wg.Add(1)
		cj.wg = wg
		cj.Number = i + 1
		jobChan <- cj
	}

	// You need wait all jobs
	wg.Wait()

	fmt.Printf("final count: %d (%v)\n", count, time.Since(start))
}

```

### TODO

- Add tests