package goworker

import (
	"sync"
)

type Queen struct {
	size  int64
	works chan func()
	wg    sync.WaitGroup
}

func NewQueen(size int64) *Queen {
	q := &Queen{
		size:  size,
		works: make(chan func(), size),
	}

	return q
}

func (q *Queen) Add(job func()) {
	q.wg.Add(1)

	go func() {
		q.works <- job
	}()
}

func (q *Queen) Run() {
	for i := 0; i < int(q.size); i++ {
		go func() {
			for f := range q.works {
				f()
				q.wg.Done()
			}
		}()
	}

	q.wg.Wait()

	close(q.works)
}
