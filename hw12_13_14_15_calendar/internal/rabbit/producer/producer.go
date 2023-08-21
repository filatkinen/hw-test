package producer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/scheduler"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/config/server"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/logger"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/rabbit"
	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/server/calendar"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	log     *logger.Logger
	config  rabbit.Config
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	app     *calendar.App
	chExit  chan struct{}
}

func NewProducer(config scheduler.Config, log *logger.Logger) (*Producer, error) {
	p := Producer{log: log, config: config.Rabbit}
	app, err := calendar.New(log, server.Config{
		LogLevel:  config.LogLevel,
		StoreType: config.StoreType,
		DB:        config.DB,
	})
	if err != nil {
		return nil, err
	}
	p.app = app

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

	queue, err := channel.QueueDeclare(
		config.Rabbit.Queue, // name of the queue
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // noWait
		nil,                 // arguments
	)
	if err != nil {
		e := p.Close()
		return nil, errors.Join(err, e)
	}
	p.queue = queue
	p.chExit = make(chan struct{})

	return &p, nil
}

func (p *Producer) Start() {
	p.log.Logging(logger.LevelInfo, "Starting Scheduller")
	timer := time.NewTicker(p.config.CheckInterval)
	defer timer.Stop()
	for {
		select {
		case <-p.chExit:
			return
		case <-timer.C:
			p.DeleteOldEvents()
			p.GetEventsAndSendMessages()
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
	if p.app != nil {
		if e := p.app.Close(context.Background()); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}

func (p *Producer) DeleteOldEvents() {
	ctx := context.Background()
	events, err := p.app.GetEventsToDelete(ctx, time.Now().UTC())
	if err != nil {
		p.log.Error("Error while getting events to delete: " + err.Error())
		return
	}
	for i := range events {
		err = p.app.DeleteEvent(ctx, events[i].ID)
		if err != nil {
			p.log.Error("Error deleting event:" + err.Error())
			continue
		}
		p.log.Logging(logger.LevelInfo,
			fmt.Sprintf("Deleted old event:%s,  Time to start:%s:", events[i].ID, events[i].DateTimeStart))
	}
}

func (p *Producer) GetEventsAndSendMessages() {
	ctx := context.Background()
	notices, err := p.app.ListNoticesToSend(ctx, time.Now().UTC())
	if err != nil {
		p.log.Error("Error while getting events to notice: " + err.Error())
		return
	}
	for i := range notices {
		b, err := rabbit.Serialize(*notices[i])
		if err != nil {
			p.log.Error("Error while serializing: " + err.Error())
			return
		}
		err = p.channel.PublishWithContext(context.Background(),
			"",           // exchange
			p.queue.Name, // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        b,
			})
		if err != nil {
			p.log.Error("Error publishing: " + err.Error())
			return
		}
		p.log.Logging(logger.LevelInfo,
			fmt.Sprintf("Sending to the rabbit notice event:%s,  Time to start:%s:", notices[i].ID, notices[i].DateTime))
	}
}
