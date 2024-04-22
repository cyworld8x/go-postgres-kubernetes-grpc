package paseto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPasetoMaker_CreateToken(t *testing.T) {
	// Create a new PasetoMaker
	maker, errMaker := NewPasetoMaker()
	assert.NoError(t, errMaker)
	// Create a token with a duration of 1 hour
	token, payload, err := maker.CreateToken("testuser", time.Hour)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotNil(t, payload)
}

func TestPasetoMaker_VerifyToken(t *testing.T) {
	// Create a new PasetoMaker
	maker, errMaker := NewPasetoMaker()
	assert.NoError(t, errMaker)

	// Create a token with a duration of 1 hour
	token, _, err := maker.CreateToken("testuser", time.Hour)
	assert.NoError(t, err)

	// Verify the token
	payload, err := maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, payload)
}
