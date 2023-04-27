package main

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

type EnvTestRecordResult struct {
	env        string
	value      string
	neadDelete bool
}
type EnvTestRecord struct {
	name   string
	data   []byte
	result EnvTestRecordResult
}

var envtest = []EnvTestRecord{
	{
		name: "BUF",
		data: []byte("\x00hello"),
		result: EnvTestRecordResult{
			env:        "BUF",
			value:      "\nhello",
			neadDelete: false,
		},
	},
	{
		name: "PROTO",
		data: []byte("123\x0045\n6789"),
		result: EnvTestRecordResult{
			env:        "PROTO",
			value:      "123\n45",
			neadDelete: false,
		},
	},
	{
		name: "FOO",
		data: []byte("   foo   \nNext String"),
		result: EnvTestRecordResult{
			env:        "FOO",
			value:      "   foo",
			neadDelete: false,
		},
	},
	{
		name: "ZERO",
		data: []byte(""),
		result: EnvTestRecordResult{
			env:        "ZERO",
			value:      "",
			neadDelete: true,
		},
	},
}

func TestReadDir(t *testing.T) {
	resultmap := make(map[string]EnvTestRecordResult, len(envtest))
	for _, v := range envtest {
		resultmap[v.name] = v.result
	}
	dir, err := os.MkdirTemp("", "env")
	if err != nil {
		log.Fatal("Cannot make directory in tmp dir")
	}
	for _, v := range envtest {
		f, err := os.Create(path.Join(dir, v.name))
		defer func() { _ = f.Close() }()
		if err != nil {
			log.Printf("Cannot make file %s in directory %s", v.name, dir)
		}
		_, err = f.Write(v.data)
		if err != nil {
			log.Printf("Cannot write data to the file %s", path.Join(dir, v.name))
		}
		err = f.Close()
		if err != nil {
			log.Printf("Cannot cloase file %s", path.Join(dir, v.name))
		}
	}

	t.Run("Check wrong directory return code", func(t *testing.T) {
		_, err := ReadDir("fake")
		require.ErrorIs(t, err, ErrOpenDirectory)
	})

	t.Run("Check some env..", func(t *testing.T) {
		env, err := ReadDir(dir)
		if err != nil {
			log.Fatalf("Cannot read env files from %s directory", dir)
		}
		require.Nil(t, err)
		require.Equal(t, len(env), len(envtest))
		for k, v := range env {
			e, ok := resultmap[k]
			require.True(t, ok)
			require.Equal(t, e.value, v.Value)
			require.Equal(t, e.neadDelete, v.NeedRemove)
		}
	})
}
