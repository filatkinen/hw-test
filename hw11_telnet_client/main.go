package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Options struct {
	host    string
	port    string
	timeout time.Duration
}

func GetFlags() (Options, error) {
	const timeout = time.Second * 10
	var options Options
	fl := flag.NewFlagSet("main", flag.ContinueOnError)
	fl.DurationVar(&options.timeout, "timeout", timeout, "timeout")
	fl.Usage = func() {
	}
	err := fl.Parse(os.Args[1:])
	if err != nil || fl.NArg() < 2 {
		fmt.Printf("Please use: %s timeout host port\n", os.Args[0])
		fmt.Printf("Where timeout is optional and default=%s, f.e: timeout=5s  or timeout=2m\n", timeout)
		return options, errors.New("bad parameters")
	}
	options.host = fl.Args()[0]
	options.port = fl.Args()[1]
	return options, nil
}

func receiveRoutine(tc TelnetClient) {
	err := tc.Receive()
	if err != nil {
		log.Printf("Got errror during receiving: %v\n", err)
	}
}

func sendRoutine(tc TelnetClient, cancelChan chan struct{}, wg *sync.WaitGroup, in *bytes.Buffer) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-cancelChan:
			return
		default:
			in.WriteString(scanner.Text() + "\n")
			err := tc.Send()
			if err != nil {
				log.Printf("Got error: %v\n", err)
				return
			}
		}
	}
	if scanner.Err() == nil {
		fmt.Printf("...EOF")
	}
}

func main() {
	in := new(bytes.Buffer)

	options, err := GetFlags()
	if err != nil {
		return
	}
	tc := NewTelnetClient(net.JoinHostPort(options.host, options.port), options.timeout, io.NopCloser(in), os.Stdout)

	err = tc.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := tc.Close()
		if err != nil {
			log.Printf("Got error while closing TelnetClient: %v\n", err)
		}
	}()

	exitSignal := make(chan os.Signal, 1)
	cancelChan := make(chan struct{})
	signal.Notify(exitSignal, syscall.SIGINT)
	go func() {
		<-exitSignal
		close(cancelChan)
	}()

	go func() {
		receiveRoutine(tc)
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		sendRoutine(tc, cancelChan, wg, in)
	}()
	wg.Wait()
}
