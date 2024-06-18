package api

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/infrastructure/repository"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/crawler"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/prompt"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/usecases/sources"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/playwright-community/playwright-go"
	"github.com/tmc/langchaingo/llms/openai"
)

type App struct {
	Server *Server
}

func New(
	// cfg *config.Config,

	Server *Server,
) *App {
	return &App{
		// Cfg:               cfg,
		Server: Server,
	}
}

func Init(
	dbConnStr postgres.DBConnString,
	pw *playwright.Playwright,
	llm *openai.LLM,
) (*App, error) {

	connPool, _ := pgxpool.New(context.Background(), string(dbConnStr))
	repository := repository.NewRepository(connPool)
	sourcesUC := sources.NewService(repository)
	crawlerUC := crawler.NewCrawlerHandler(repository, pw)
	promptUC := prompt.NewPromptHandler(llm)
	apiServer := NewServer(sourcesUC, crawlerUC, promptUC)
	app := New(apiServer)
	return app, nil

}
