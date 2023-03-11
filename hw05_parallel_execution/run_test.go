package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		_, err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		_, err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)
		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestRunAdditional(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("Check if ErrorLimit <0. So we take it as 0.  ", func(t *testing.T) {
		sCount := 50
		eCount := 50
		tasksCount := sCount + eCount
		workersCount := 10
		maxErrorsCount := -1
		tasks := make([]Task, 0, tasksCount)

		for i := 0; i < sCount; i++ {
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				return nil
			})
		}
		for i := sCount; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				return err
			})
		}
		result, err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, result.countSuccess+result.countError, sCount+workersCount, "extra tasks were started")
	})

	t.Run("Check concurrency if sleeptime in function==0, functions=1000, workers=100", func(t *testing.T) {
		tasksCount := 1_000
		workersCount := 100
		maxErrorsCount := 0
		tasks := make([]Task, 0, tasksCount)

		for i := 0; i < tasksCount; i++ {
			tasks = append(tasks, func() error {
				return nil
			})
		}
		result, err := Run(tasks, workersCount, maxErrorsCount)

		require.Nil(t, err)
		require.Equal(t, result.countSuccess+result.countError, tasksCount)
	})

	t.Run("Checking  if workers<1", func(t *testing.T) {
		tasksCount := 100
		workersCount := 0
		maxErrorsCount := 1
		tasks := make([]Task, 0, tasksCount)

		for i := 0; i < tasksCount; i++ {
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				return nil
			})
		}
		_, err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsWorkerHasBePositive), "actual err - %v", err)
	})
}

func TestRunAdditionalTimer(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("Checking concurrency time execution-avoid sleep CI issue  - with Ticker", func(t *testing.T) {
		sCount := 50
		eCount := 0
		tasksCount := sCount + eCount
		workersCount := 10
		maxErrorsCount := 0
		tasks := make([]Task, 0, tasksCount)
		durationFunc := time.Millisecond * 100
		durationTicker := time.Millisecond * 2

		for i := 0; i < sCount; i++ {
			tasks = append(tasks, func() error {
				startTime := time.Now()
				ticker := time.NewTicker(durationTicker)
				for range ticker.C {
					if time.Since(startTime) >= durationFunc {
						break
					}
				}
				ticker.Stop()
				return nil
			})
		}

		start := time.Now()
		result, err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := int64(time.Since(start))
		estimatedTime := int64(durationFunc) * int64(tasksCount) / int64(workersCount)
		ratio := float32(elapsedTime) / float32(estimatedTime)

		require.Nil(t, err)
		require.Equal(t, result.countSuccess+result.countError, tasksCount)
		fmt.Printf("Ratio of elapsed to estimated test time: %.2f\n", ratio)
		require.LessOrEqual(t, ratio, float32(1.2), "tasks were run sequentially?")
	})
}
