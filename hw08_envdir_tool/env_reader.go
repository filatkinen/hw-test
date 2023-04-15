package main

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var ErrOpenDirectory = errors.New("can not read directory")

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func readEnvFile(fileName string) (e string, err error) {
	fenv, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer fenv.Close()
	scanner := bufio.NewScanner(fenv)
	scanner.Scan()
	b := bytes.ReplaceAll(scanner.Bytes(), []byte("\x00"), []byte("\n"))
	return strings.TrimRightFunc(string(b), unicode.IsSpace), nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(map[string]EnvValue, 10)
	dirRead, err := os.Open(dir)
	if err != nil {
		return nil, ErrOpenDirectory
	}
	defer dirRead.Close()

	files, err := dirRead.ReadDir(0)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fname := filepath.Join(dir, file.Name())
		f, err := file.Info()
		if err != nil {
			log.Printf("Can not get stat information of file %s, miss it", fname)
			continue
		}
		if strings.Contains(f.Name(), "=") {
			log.Printf("File %s contain in it's name =, miss it", fname)
			continue
		}

		e := EnvValue{Value: "", NeedRemove: true}

		if f.Size() != 0 {
			e.NeedRemove = false
			e.Value, err = readEnvFile(fname)
			if err != nil {
				log.Printf("Can not read or open file %s, miss it", fname)
				continue
			}
			if len(e.Value) == 0 {
				e.NeedRemove = true
			}
		}
		env[f.Name()] = e
	}
	return env, nil
}
