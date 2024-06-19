package crawler

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/crawler/infrastructure/repository"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

type service struct {
	repo repository.Repository
	pw   *playwright.Playwright
}

func NewService(repo repository.Repository, pw *playwright.Playwright) UseCase {
	return &service{
		repo: repo,
		pw:   pw,
	}
}

func (s *service) Get(ctx context.Context, arg *domain.WebSite) (*domain.Entry, error) {
	var collection domain.Entry = make(domain.Entry)
	browser, err := s.pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})

	if err != nil {
		log.Error().Err(err).Msgf("could not launch browser: %v", err)
	} //else {
	// 	defer func() {
	// 		if err := browser.Close(); err != nil {
	// 			log.Error().Err(err).Msgf("could not close browser: %v", err)
	// 		}
	// 	}()
	// }
	//TODO: Fix the defer close browser

	// for i := 0; i < 10; i++ {
	// 	go func() {
	newPage, err := browser.NewPage()
	if err != nil {
		log.Error().Err(err).Msgf("could not create page: %v", err)
	}

	lo.ForEach(arg.Pages, func(page domain.Page, _ int) {

		if _, err = newPage.Goto(page.Url); err != nil {
			log.Error().Err(err).Msgf("could not goto: %v. Url: %s", err, page.Url)
		}

		lo.ForEach(page.PageEvents, func(pageEvent domain.PageEvent, _ int) {
			PageDo(pageEvent, newPage)

			if len(pageEvent.Collectors) > 0 {
				lo.ForEach(pageEvent.Collectors, func(pageObject *domain.PageObject, _ int) {
					Parse(pageObject, newPage, collection)
				})
			}
		})
	})
	time.Sleep(10 * time.Second)
	newPage.Close()
	// 	}()
	// }

	return &collection, nil
}

func PageDo(pageEvent domain.PageEvent, newPage playwright.Page) {
	if pageEvent.Selector != "" {
		var eventErr error
		switch strings.ToLower(pageEvent.Type) {
		case "fill":
			eventErr = newPage.Locator(pageEvent.Selector).Last().Fill(pageEvent.EnterValue)
		case "click":
			eventErr = newPage.Locator(pageEvent.Selector).Last().Click()
		}
		if eventErr != nil {
			log.Error().Err(eventErr).Msgf("could not perform event: %v, error: %v", pageEvent.Selector, eventErr)
		}
	}
}

func Parse(pageObject *domain.PageObject, newPage playwright.Page, collection domain.Entry) {
	if len(pageObject.PageObject) > 0 {
		var contents domain.Entry = make(domain.Entry)
		log.Info().Msgf("Parsing page object: %v", pageObject.Key)
		lo.ForEach(pageObject.PageObject, func(subPageObject *domain.PageObject, index int) {
			subContents, errs := GetContent(subPageObject, newPage)
			LogErrors(errs)
			if subContents != nil {
				contents[subPageObject.Key] = subContents
			}
		})
		collection[pageObject.Key] = contents
	} else {
		content, errs := GetContent(pageObject, newPage)
		if len(errs) > 0 {
			for _, err := range errs {
				log.Error().Err(err).Msgf("could not get page object: %v", err)
			}
		}
		collection[pageObject.Key] = content
	}
}

func GetContent(pageObject *domain.PageObject, page playwright.Page) (*string, []error) {
	var listErrors []error = make([]error, 0)
	selectors, err := page.Locator(pageObject.Selector).All()
	if err != nil {
		listErrors = append(listErrors, err)
	}
	if len(selectors) > 0 {
		var textContents []string
		lo.ForEach(selectors, func(selector playwright.Locator, _ int) {
			text, errGetText := selector.TextContent()
			if err != nil {
				listErrors = append(listErrors, errGetText)
			}
			textContents = append(textContents, text)
		})
		if textContents != nil {
			content := strings.Join(textContents, ",")
			return &content, listErrors
		}
	} else {
		listErrors = append(listErrors, errors.New("no selectors found "+pageObject.Selector))
	}
	return nil, listErrors
}

func LogErrors(errs []error) {
	if len(errs) > 0 {
		for _, err := range errs {
			log.Error().Err(err).Msgf("could not get page object: %v", err)
		}
	}
}
