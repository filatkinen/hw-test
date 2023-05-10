package main

import (
	"bufio"
	"errors"
	"io"
	"net"
	"time"
)

var (
	ErrErrorsConnect = errors.New("unable to connect")
	ErrErrorsSend    = errors.New("unable to send")
	ErrErrorsReceive = errors.New("unable to receive")
	ErrErrorsClose   = errors.New("error with close")
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct {
	in      io.ReadCloser
	out     io.Writer
	timeout time.Duration
	address string
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Telnet{
		in:      in,
		out:     out,
		timeout: timeout,
		address: address,
	}
}

func (t *Telnet) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return errors.Join(ErrErrorsConnect, err)
	}
	t.conn = conn
	return nil
}

func (t *Telnet) Send() error {
	scanner := bufio.NewScanner(t.in)
	for scanner.Scan() {
		_, err := t.conn.Write(append(scanner.Bytes(), '\n'))
		if err != nil {
			return errors.Join(ErrErrorsSend, err)
		}
	}
	if scanner.Err() != nil {
		return errors.Join(ErrErrorsSend, scanner.Err())
	}
	return nil
}

func (t *Telnet) Receive() error {
	scanner := bufio.NewScanner(t.conn)
	for scanner.Scan() {
		_, err := t.out.Write(append(scanner.Bytes(), '\n'))
		if err != nil {
			return errors.Join(ErrErrorsReceive, err)
		}
	}
	if scanner.Err() != nil {
		return errors.Join(ErrErrorsReceive, scanner.Err())
	}
	return nil
}

func (t *Telnet) Close() error {
	if t.conn == nil {
		return errors.Join(ErrErrorsClose, errors.New("connection is already closed"))
	}
	err := t.conn.Close()
	if err != nil {
		return errors.Join(ErrErrorsClose, err)
	}
	return nil
}
