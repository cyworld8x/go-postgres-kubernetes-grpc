package paseto

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Issuer    string    `json:"issuer,omitempty"`
	Audience  string    `json:"audience,omitempty"`
	ExpiredAt time.Time `json:"expired_at ,omitempty"`
	IssuedAt  time.Time `json:"issued_at,omitempty"`
}

func (p *Payload) IsValid() bool {
	now := time.Now().UTC()

	if now.Before(p.IssuedAt) {
		return false
	}

	if now.After(p.ExpiredAt) {
		return false
	}

	return true
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	now := time.Now().UTC()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:        tokenID,
		Issuer:    "issuer",
		Audience:  "audience",
		Username:  username,
		ExpiredAt: now.Add(duration),
		IssuedAt:  now,
	}, nil
}
