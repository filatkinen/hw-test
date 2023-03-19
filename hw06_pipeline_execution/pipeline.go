package hw06pipelineexecution

import (
	"sync"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	sm := sync.Map{}
	wgPipeLine := sync.WaitGroup{}
	outChan := make(Bi)
	finishChan := make(chan struct{})

	// Pipeline function
	fstage := func(val any, i int) {
		defer wgPipeLine.Done()
		for _, stage := range stages {
			select {
			case <-done:
				return // Do not send unfinished result
			default:
				chIn := make(Bi)
				go func() { chIn <- val }()
				val = <-stage(chIn)
				close(chIn)
			}
		}
		sm.Store(i, val)
	}

	// Sender
	go func() {
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			default:
				select {
				case <-done:
					return
				case val, ok := <-in:
					if ok {
						wgPipeLine.Add(1)
						go fstage(val, i)
					} else {
						close(finishChan)
						return
					}
				}
			}
		}
	}()

	// Watcher for ending  Pipeline
	go func() {
		select {
		case <-finishChan:
			wgPipeLine.Wait()
			for i := 0; ; i++ {
				val, ok := sm.Load(i)
				if !ok {
					break
				}
				outChan <- val
			}
		case <-done:
			break
		}
		close(outChan)
	}()

	return outChan
}
