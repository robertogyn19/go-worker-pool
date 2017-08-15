package worker_pool

// Worker struct
type Worker struct {
	WorkerPool chan chan GenericJob
	JobChannel chan GenericJob
	quit       chan bool
}

// NewWorker func
func NewWorker(workerPool chan chan GenericJob) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan GenericJob),
		quit:       make(chan bool),
	}
}

// Start starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start(id int) {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				job.Start(id)
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}
