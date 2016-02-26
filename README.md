### Worker-Pool

A "better" version of this [example](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/)

Simple and useful (for me) worker-pool with configurable queue size.
Usage example:

``` golang
package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"

	wp "github.com/robertogyn19/go-worker-pool"
)

type CustomJob struct {
	Number int
	wg *sync.WaitGroup
}

func (cj *CustomJob) Start() {
	work(cj.Number)
	cj.wg.Done()
}

var count int

func work(i int) {
	r := rand.Intn(100)
	d := time.Duration(r) * time.Millisecond
	fmt.Printf("(%d) Sleeping for %v\n", i, d)
	time.Sleep(d)
	count += 1
}

func main() {
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

	fmt.Println("Final count:", count)
}
```

### TODO

- Add tests