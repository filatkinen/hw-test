package rabbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/filatkinen/hw-test/hw12_13_14_15_calendar/internal/storage"
)

const (
	DefaultChecKTimeSheduler time.Duration = time.Second * 5
)

type Config struct {
	Port          string
	Address       string
	User          string
	Password      string
	Queue         string
	CheckInterval time.Duration
	Tag           string
}

func (r Config) GetDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", r.User, r.Password, r.Address, r.Port)
}

func Serialize(msg storage.Notice) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func Deserialize(b []byte) (storage.Notice, error) {
	var msg storage.Notice
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
