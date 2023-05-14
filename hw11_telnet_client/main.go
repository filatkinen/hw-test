package main

import (
	"errors"
	"flag"
	"fmt"
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

func main() {
	options, err := GetFlags()
	if err != nil {
		return
	}
	tc := NewTelnetClient(net.JoinHostPort(options.host, options.port), options.timeout, os.Stdin, os.Stdout)

	err = tc.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = tc.Close()
	}()

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT)
	go func() {
		<-exitSignal
		err := tc.Close()
		if err != nil {
			log.Printf("Got error during closing: %v\n", err)
		}
	}()

	go func() {
		err := tc.Receive()
		if err != nil {
			log.Printf("Got error during receiving: %v\n", err)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := tc.Send()
		if err != nil {
			log.Printf("Got error during sending: %v\n", err)
			return
		}
		fmt.Println("...EOF")
	}()
	wg.Wait()
}
