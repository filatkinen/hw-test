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

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (resultFunc ResultTask, err error) {
	if n < 1 {
		return resultFunc, ErrErrorsWorkerHasBePositive
	}
	if m < 0 {
		m = 0
	}
	countError := 0
	countSuccess := 0
	cancelChan := make(chan struct{})
	resultChan := make(chan error, n)

	// Send tasks to chan  with buffer length=min(len(tasks),n)
	taskBufferSize := len(tasks)
	if taskBufferSize > n {
		taskBufferSize = n
	}
	countSendTask := taskBufferSize
	taskChan := make(chan Task, taskBufferSize)
	for i := 0; i < taskBufferSize; i++ {
		taskChan <- tasks[i]
	}

	// Starting "n" handlers
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case t := <-taskChan:
					resultChan <- t()
				case <-cancelChan:
					return
				}
			}
		}()
	}

	funcCountResult := func(result error) {
		if result != nil {
			if err != nil {
				err = fmt.Errorf(" %w", result)
			} else {
				err = result
			}
			countError++
		} else {
			countSuccess++
		}
	}

	for result := range resultChan {
		funcCountResult(result)
		tasksLeft := len(tasks) - countSendTask
		if !(countError > m) && tasksLeft > 0 {
			countSendTask++
			taskChan <- tasks[countSendTask-1]
		} else {
			close(cancelChan)
			break
		}
	}

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		funcCountResult(result)
	}

	if countError > m {
		err = fmt.Errorf(" %w", ErrErrorsLimitExceeded)
	}
	resultFunc.countError = countError
	resultFunc.countSuccess = countSuccess

	return resultFunc, err
}
