package sessions

import (
	"context"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
)

// Service struct
type Service struct {
	sessionRepository domain.SessionRepository
}

// NewService creates a new session service
func NewService(sessionRepository domain.SessionRepository) *Service {
	return &Service{
		sessionRepository: sessionRepository,
	}
}

// GetSession retrieves a session by its ID
func (s *Service) GetSession(ctx context.Context, sessionID string) (*domain.Session, error) {
	return s.sessionRepository.GetSession(ctx, sessionID)
}

// DeleteSession deletes a session by its ID
func (s *Service) DeleteSession(ctx context.Context, sessionID string) error {
	return s.sessionRepository.DeleteSession(ctx, sessionID)
}

// UpdateSession updates an existing session
func (s *Service) UpdateSession(ctx context.Context, session *domain.Session) error {
	return s.sessionRepository.UpdateSession(ctx, session)
}

// ListSessions lists all sessions
func (s *Service) ListSessions(ctx context.Context) ([]*domain.Session, error) {
	return s.sessionRepository.ListSessions(ctx)
}

// GetSessionByUserID retrieves sessions by user ID
func (s *Service) GetSessionByUserID(ctx context.Context, username string) ([]*domain.Session, error) {
	return s.sessionRepository.GetSessionByUserID(ctx, username)
}

// BlockSession blocks a session by its ID
func (s *Service) BlockSession(ctx context.Context, sessionID string) (bool, error) {
	return s.sessionRepository.BlockSession(ctx, sessionID)
}

// GenerateSession generates a new session for a user
func (s *Service) GenerateSession(ctx context.Context, username string, token string, duration time.Duration) (*domain.Session, error) {
	return s.sessionRepository.GenerateSession(ctx, username, token, duration)
}

// UnblockSession unblocks a session by its ID
func (s *Service) UnblockSession(ctx context.Context, sessionID string) (bool, error) {
	return s.sessionRepository.UnblockSession(ctx, sessionID)
}

// ClearSession clears a session by its token
func (s *Service) ClearSession(ctx context.Context, token string, username string) error {
	return s.sessionRepository.ClearSession(ctx, token, username)
}
