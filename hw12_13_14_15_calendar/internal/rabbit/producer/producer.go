package producer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/scheduler"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	log     *logger.Logger
	config  rabbit.Config
	conn    *amqp.Connection
	channel *amqp.Channel
	chExit  chan struct{}
}

func NewProducer(config scheduler.Config, log *logger.Logger) (*Producer, error) {
	p := Producer{log: log, config: config.Rabbit}

	configAmpq := amqp.Config{
		Vhost:      "/",
		Properties: amqp.NewConnectionProperties(),
	}
	configAmpq.Properties.SetClientConnectionName("scheduller")

	conn, err := amqp.Dial(config.Rabbit.GetDSN())
	if err != nil {
		e := p.Close()
		return nil, errors.Join(err, e)
	}
	p.conn = conn

	channel, err := conn.Channel()
	if err != nil {
		e := p.Close()
		return nil, errors.Join(err, e)
	}
	p.channel = channel

	err = channel.ExchangeDeclare(
		config.Rabbit.ExchangeName, // name
		"fanout",                   // type
		true,                       // durable
		false,                      // auto-deleted
		false,                      // internal
		false,                      // no-wait
		nil,                        // arguments
	)
	if err != nil {
		e := p.Close()
		return nil, errors.Join(err, e)
	}
	p.channel = channel

	p.chExit = make(chan struct{})

	return &p, nil
}

func (p *Producer) Start(f func() [][]byte) {
	p.log.Logging(logger.LevelInfo, "Starting Scheduller")
	timer := time.NewTicker(p.config.CheckInterval)
	defer timer.Stop()
	for {
		select {
		case <-p.chExit:
			return
		case <-timer.C:
			p.SendMessages(f())
		}
	}
}

func (p *Producer) Stop() {
	p.log.Logging(logger.LevelInfo, "Stopping Scheduller")
	p.chExit <- struct{}{}
}

func (p *Producer) Close() (err error) {
	p.log.Logging(logger.LevelInfo, "Closing Scheduller")
	if p.channel != nil {
		if e := p.channel.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}
	if p.conn != nil {
		if e := p.conn.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}

func (p *Producer) SendMessages(messages [][]byte) {
	for i := range messages {
		err := p.channel.PublishWithContext(context.Background(),
			p.config.ExchangeName, // exchange
			"",                    // routing key
			false,                 // mandatory
			false,                 // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        messages[i],
			})
		if err != nil {
			p.log.Error("Error publishing: " + err.Error())
			return
		}
		var b []byte
		if len(messages[i]) < 60 {
			b = messages[i][:len(messages[i])]
		} else {
			b = messages[i][:60]
		}
		p.log.Logging(logger.LevelInfo,
			fmt.Sprintf("Sending to the rabbit notice event. First 60 symbols of message:%s:", string(b)))
	}
}
