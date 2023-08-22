package consumer

import (
	"errors"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/sender"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Comsumer struct {
	log     *logger.Logger
	config  rabbit.Config
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	chExit  chan struct{}
	chMsg   <-chan amqp.Delivery
}

func NewConsumer(config sender.Config, log *logger.Logger) (*Comsumer, error) {
	c := Comsumer{log: log, config: config.Rabbit}

	configAmpq := amqp.Config{
		Vhost:      "/",
		Properties: amqp.NewConnectionProperties(),
	}
	configAmpq.Properties.SetClientConnectionName("noticer")

	conn, err := amqp.Dial(config.Rabbit.GetDSN())
	if err != nil {
		e := c.Close()
		return nil, errors.Join(err, e)
	}
	c.conn = conn

	channel, err := conn.Channel()
	if err != nil {
		e := c.Close()
		return nil, errors.Join(err, e)
	}
	c.channel = channel

	queue, err := channel.QueueDeclare(
		config.Rabbit.Queue, // name of the queue
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // noWait
		nil,                 // arguments
	)
	if err != nil {
		e := c.Close()
		return nil, errors.Join(err, e)
	}
	c.queue = queue
	c.chExit = make(chan struct{})

	msgs, err := c.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		e := c.Close()
		return nil, errors.Join(err, e)
	}
	c.chMsg = msgs
	return &c, nil
}

func (c *Comsumer) Start(f func([]byte)) {
	c.log.Logging(logger.LevelInfo, "Starting Noticer")
	go func() {
		for d := range c.chMsg {
			f(d.Body)
		}
	}()
	<-c.chExit
}

func (c *Comsumer) Stop() {
	c.log.Logging(logger.LevelInfo, "Stopping Noticer")
	c.chExit <- struct{}{}
}

func (c *Comsumer) Close() (err error) {
	c.log.Logging(logger.LevelInfo, "Closing Noticer")
	if c.channel != nil {
		if e := c.channel.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}
	if c.conn != nil {
		if e := c.conn.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}
