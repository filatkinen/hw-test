package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type CmdRecord struct {
	comargs []string
	code    int
}

var cmd = []CmdRecord{
	{
		comargs: []string{"ls", "-l"},
		code:    0,
	},
	{
		comargs: []string{"ls", "/fakedirectory"},
		code:    2,
	},
}

func TestRunCmd(t *testing.T) {
	t.Run("Check results codes", func(t *testing.T) {
		for _, v := range cmd {
			code := RunCmd(v.comargs, Environment{})
			require.Equal(t, v.code, code, v)
		}
	})
}
