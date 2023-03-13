package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrErrorsLimitExceeded       = errors.New("errors limit exceeded")
	ErrErrorsWorkerHasBePositive = errors.New("количество обработчиков не может быть меньше 1")
)

type Task func() error

type ResultTask struct {
	countSuccess int
	countError   int
}

func runWorker(task <-chan Task, cancel <-chan struct{}, resul chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-cancel:
			return
		default:
			select {
			case t, ok := <-task:
				if ok {
					resul <- t()
				} else {
					return
				}
			case <-cancel:
				return
			}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (result ResultTask, err error) {
	if n < 1 {
		return result, ErrErrorsWorkerHasBePositive
	}
	if m < 0 {
		m = 0
	}

	cancelChanWorkers := make(chan struct{})
	resultChan := make(chan error)
	taskChan := make(chan Task, len(tasks))

	// Starting "n" handlers
	wgWorker := sync.WaitGroup{}
	wgWorker.Add(n)
	for i := 0; i < n; i++ {
		go runWorker(taskChan, cancelChanWorkers, resultChan, &wgWorker)
	}

	// Send all the tasks to the chanel for workers.
	// They send back result to the channel without buffer, so we control them by result channel.
	for i := 0; i < len(tasks); i++ {
		taskChan <- tasks[i]
	}
	close(taskChan)

	// Starting resulting worker to receive results
	wgResult := sync.WaitGroup{}
	wgResult.Add(1)
	go func() {
		defer wgResult.Done()
		once := sync.Once{}
		for r := range resultChan {
			if r != nil {
				err = fmt.Errorf(" %w", r)
				result.countError++
				if result.countError >= m {
					once.Do(func() { close(cancelChanWorkers) })
				}
			} else {
				result.countSuccess++
			}
		}
	}()

	// Waiting till all the workers closed.
	wgWorker.Wait()
	// Close resultChan to stop result worker.
	close(resultChan)
	// Waiting till result worker closed.
	wgResult.Wait()

	if result.countError > m {
		err = fmt.Errorf(" %w", ErrErrorsLimitExceeded)
	}
	return result, err
}
