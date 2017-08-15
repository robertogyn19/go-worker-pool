package worker_pool

/* -------------------------------------------------------------------------- */
/*																																						*/
/*												         Dispatcher        													*/
/*																																						*/
/* -------------------------------------------------------------------------- */

// Dispatcher struct
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan GenericJob
	MaxWorkers int
	Workers    []Worker
	JobQueue   chan GenericJob
}

// NewDispatcher func
func NewDispatcher(maxWorkers int, jobQueue chan GenericJob) *Dispatcher {
	pool := make(chan chan GenericJob, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers, JobQueue: jobQueue}
}

func (d *Dispatcher) StopWorkers() {
	for _, w := range d.Workers {
		w.quit <- true
	}
}

// Run func
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start(i+1)

		d.Workers = append(d.Workers, worker)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			// a job request has been received
			go func(job GenericJob) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
