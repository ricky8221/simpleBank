package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssueAt   time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expire_at"`
}

// NewPayload creates a new token with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid check if the token payload is valid or not
func (Payload *Payload) Valid() error {
	if time.Now().After(Payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
