package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(10)
		for i := 0; i < 10; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}

		c = c.Clear()

		for i := 0; i < 10; i++ {
			_, ok := c.Get(Key(strconv.Itoa(i)))
			require.False(t, ok)
		}

		require.Equal(t, 0, c.Len())
	})
}

func TestCacheAdditional(t *testing.T) {
	t.Run(`Logic push out due size of 3 elements: add 4 elements, first go out `, func(t *testing.T) {
		c := NewCache(3)

		sliceToPush := []struct {
			key   string
			value int
		}{
			{key: "aaa", value: 10},
			{key: "bbb", value: 20},
			{key: "ccc", value: 30},
			{key: "ddd", value: 40},
		}

		for _, v := range sliceToPush {
			c.Set(Key(v.key), v.value)
		}
		for idx, v := range sliceToPush {
			if idx == 0 {
				_, ok := c.Get(Key(v.key))
				require.False(t, ok)
				continue
			}
			val, ok := c.Get(Key(v.key))
			require.True(t, ok)
			require.Equal(t, val, v.value)
		}
	})

	t.Run(`Logic push out the oldest accessed element`, func(t *testing.T) {
		c := NewCache(3)

		sliceToPush := []struct {
			key   string
			value int
		}{
			{key: "aaa", value: 10},
			{key: "bbb", value: 20},
			{key: "ccc", value: 30}, // candidate to push out
		}

		sliceToCheck := []struct {
			key   string
			value int
		}{
			{key: "aaa", value: 10},
			{key: "bbb", value: 20},
			{key: "ddd", value: 40},
		}

		for _, v := range sliceToPush {
			c.Set(Key(v.key), v.value)
		}

		c.Get("aaa")
		c.Get("bbb")
		c.Set("aaa", 10)
		c.Get("aaa")
		c.Set("ddd", 40)
		for _, v := range sliceToCheck {
			val, ok := c.Get(Key(v.key))
			require.True(t, ok)
			require.Equal(t, val, v.value)
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
