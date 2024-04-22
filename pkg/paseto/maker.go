package paseto

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type PasetoMaker struct {
	secretKey []byte
}

func NewPasetoMaker() (*PasetoMaker, error) {
	symmetricKey := "YELLOW SUBMARINE, BLACK WIZARDRY"
	if len(symmetricKey) != 32 {
		return nil, fmt.Errorf("invalid key size: must be exactly 32 characters")
	}
	return &PasetoMaker{
		secretKey: []byte(symmetricKey),
	}, nil
}

func (pm *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	// Encrypt and sign the token
	encrypter := paseto.NewV2()
	token, err := encrypter.Encrypt(pm.secretKey, payload, nil)
	if err != nil {
		return "", payload, ErrInvalidToken
	}

	return token, payload, nil
}

func (pm *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	// Create a new Paseto token
	payload := &Payload{}

	// Decrypt and verify the token
	decrypter := paseto.NewV2()
	err := decrypter.Decrypt(token, pm.secretKey, payload, nil)
	if err != nil {
		return nil, err
	}

	if payload.IsValid() {
		return payload, nil
	}
	return nil, ErrExpiredToken

}
