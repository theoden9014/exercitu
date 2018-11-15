package client

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/theoden9014/exercitu/worker"
)

type Client struct {
	worker    worker.Worker
	workers   int
	rate      time.Duration
	ticker    <-chan time.Time
	semaphore chan int
	wg        sync.WaitGroup
	duration  time.Duration
	count     uint64
	results   chan *worker.Result
	logger    *log.Logger
}

const (
	DefaultWorkers  = 5
	DefaultDuration = 0
	DefaultRate     = 0
)

func New() *Client {
	e := &Client{
		workers:  DefaultWorkers,
		rate:     DefaultRate,
		duration: DefaultDuration,
	}

	e.semaphore = make(chan int, e.workers)
	e.results = make(chan *worker.Result, e.workers)
	e.ticker = time.Tick(e.rate)

	return e
}

func (c *Client) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := c.worker.Try(ctx); err != nil {
		return err
	}

	go func() {
		defer close(c.semaphore)
		defer c.wg.Wait()

		for {
			select {
			case <-c.ticker:
				go c.run(ctx)
			}
		}
	}()

	for {
		select {
		case _, ok := <-c.results:
			if !ok {
				return nil
			}
		}
	}
}

func (c *Client) run(ctx context.Context) {
	defer c.wg.Done()

	c.wg.Add(1)
	ret, err := c.worker.Run(ctx)
	if err != nil {
		// c.logger.Error(err)
	}

	c.results <- ret
	c.semaphore <- 1
}

func (c *Client) SetWorker(w worker.Worker) {
	c.worker = w
}
