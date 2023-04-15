package hw06pipelineexecution

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	sleepPerStage = time.Millisecond * 100
	fault         = sleepPerStage / 2
)

func TestPipeline(t *testing.T) {
	// Stage generator
	g := func(_ string, f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage)
					out <- f(v)
				}
			}()
			return out
		}
	}

	stages := []Stage{
		g("Dummy", func(v interface{}) interface{} { return v }),
		g("Multiplier (* 2)", func(v interface{}) interface{} { return v.(int) * 2 }),
		g("Adder (+ 100)", func(v interface{}) interface{} { return v.(int) + 100 }),
		g("Stringifier", func(v interface{}) interface{} { return strconv.Itoa(v.(int)) }),
	}

	t.Run("simple case", func(t *testing.T) {
		in := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Equal(t, []string{"102", "104", "106", "108", "110"}, result)
		require.Less(t,
			int64(elapsed),
			// ~0.8s for processing 5 values in 4 stages (100ms every) concurrently
			int64(sleepPerStage)*int64(len(stages)+len(data)-1)+int64(fault))
	})

	t.Run("done case", func(t *testing.T) {
		in := make(Bi)
		done := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		// Abort after 200ms
		abortDur := sleepPerStage * 2
		go func() {
			<-time.After(abortDur)
			close(done)
		}()

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, done, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Len(t, result, 0)
		require.Less(t, int64(elapsed), int64(abortDur)+int64(fault))
	})
}

func TestPipelineAdd(t *testing.T) {
	// Stage generator
	g := func(_ string, delay time.Duration, f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					if delay > 0 {
						time.Sleep(delay)
					}
					out <- f(v)
				}
			}()
			return out
		}
	}

	t.Run("Stress 1000 values for 1000 functions with delay=0", func(t *testing.T) {
		in := make(Bi)

		funcInPipeLine := 1000
		dataArray := 1000

		stages := make([]Stage, 0, funcInPipeLine)
		etalonData := make([]int, 0, dataArray)
		for i := 0; i < funcInPipeLine; i++ {
			if i%2 == 0 {
				stages = append(stages, g("Adder (+3)", time.Millisecond*0, func(v interface{}) interface{} { return v.(int) + 3 }))
			} else {
				stages = append(stages, g("Adder (-2)", time.Millisecond*0, func(v interface{}) interface{} { return v.(int) - 2 }))
			}
		}

		go func() {
			for i := 0; i < dataArray; i++ {
				etalonData = append(etalonData, i+500)
				in <- i
			}
			close(in)
		}()

		result := make([]int, 0, dataArray)
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s.(int))
		}
		require.Equal(t, etalonData, result)
	})

	t.Run("Time case 100 values 10 functions 100 msec. We are expecting less then 2 seconds ", func(t *testing.T) {
		in := make(Bi)

		funcInPipeLine := 10
		dataArray := 100
		delay := time.Millisecond * 100

		stages := make([]Stage, 0, funcInPipeLine)
		for i := 0; i < funcInPipeLine; i++ {
			stages = append(stages, g("Dummy", delay, func(v interface{}) interface{} { return v }))
		}

		go func() {
			for i := 0; i < dataArray; i++ {
				in <- i
			}
			close(in)
		}()

		start := time.Now()
		for s := range ExecutePipeline(in, nil, stages...) {
			_ = s
		}
		elapsed := time.Since(start)
		require.Less(t, int64(elapsed), int64(delay)*int64(funcInPipeLine)*2)
	})
}
