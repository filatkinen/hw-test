package storage

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

type Event struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	DateTimeStart  time.Time `json:"dateTimeStart"`
	DateTimeEnd    time.Time `json:"dateTimeEnd"`
	DateTimeNotice time.Time `json:"dateTimeNotice"`
	UserID         string    `json:"userid"`
}

type User struct {
	UserID string `json:"userid"`
	Email  string `json:"email"`
}

type Notice struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	DateTime time.Time `json:"datetime"`
	UserID   string    `json:"userid"`
}

func PseudoUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return strings.ToLower(fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]))
}
