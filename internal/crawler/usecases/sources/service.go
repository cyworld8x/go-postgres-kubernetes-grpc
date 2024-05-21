package sources

import (
	"context"
	"encoding/json"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/infrastructure/repository"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/infrastructure/repository/postgres"
	_ "github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) UseCase {
	return &service{
		repo: repo,
	}
}

// NewService creates a new service
func (s *service) CreateSource(ctx context.Context, arg *domain.Source) (*domain.Source, error) {

	data, err := json.Marshal(arg.Data)
	if err != nil {
		return nil, err
	}

	createSourceParams := postgres.CreateSourceParams{Data: data, Name: arg.Name}

	source, err := s.repo.CreateSource(ctx, createSourceParams)
	if err != nil {
		log.Error().Err(err).Msg("cannot create source")
		return nil, err
	}

	website := domain.WebSite{}

	err = json.Unmarshal(source.Data, &website)

	if err != nil {
		log.Error().Err(err).Msg("cannot parse source data")
		return nil, err
	}

	return &domain.Source{
		Id:   source.ID,
		Name: source.Name,
		Data: website,
	}, nil

}
func (s *service) GetSource(ctx context.Context, id uuid.UUID) (*domain.Source, error) {
	source, err := s.repo.GetSource(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("cannot get source:" + id.String())
		return nil, err
	}
	website := domain.WebSite{}

	err = json.Unmarshal(source.Data, &website)

	if err != nil {
		log.Error().Err(err).Msg("cannot parse source data")
		return nil, err
	}

	return &domain.Source{
		Id:   source.ID,
		Name: source.Name,
		Data: website,
	}, nil

}

func (s *service) GetSources(ctx context.Context) ([]domain.Source, error) {
	sources, err := s.repo.GetSources(ctx)
	if err != nil {
		log.Error().Err(err).Msg("cannot get sources")
		return []domain.Source{}, err
	}
	var domainSources []domain.Source
	for _, source := range sources {
		website := domain.WebSite{}

		err = json.Unmarshal(source.Data, &website)

		if err != nil {
			return []domain.Source{}, err
		}
		domainSources = append(domainSources, domain.Source{
			Id:   source.ID,
			Name: source.Name,
			Data: website,
		})
	}
	return domainSources, nil
}

func (s *service) UpdateSource(ctx context.Context, arg *domain.Source) (*domain.Source, error) {
	data, err := json.Marshal(arg.Data)
	if err != nil {
		return nil, err
	}
	updateSourceParams := postgres.UpdateSourceParams{Data: data, ID: arg.Id, Name: arg.Name}
	source, err := s.repo.UpdateSource(ctx, updateSourceParams)
	if err != nil {
		log.Error().Err(err).Msg("cannot update source:" + arg.Id.String())
		return nil, err
	}
	website := domain.WebSite{}

	err = json.Unmarshal(source.Data, &website)

	if err != nil {
		return nil, err
	}
	return &domain.Source{
		Id:   source.ID,
		Name: source.Name,
		Data: website,
	}, nil
}
