package domain

import (
	"github.com/google/uuid"
)

type WebSite struct {
	Url   string `json:"url"`
	Pages []Page `json:"pages"`
}

type Page struct {
	Url        string      `json:"url"`
	PageEvents []PageEvent `json:"page_events"`
}

type PageEvent struct {
	Selector      string        `json:"selector"`
	Type          string        `json:"type"`
	EnterValue    string        `json:"enter_value"`
	TimeSleep     int           `json:"time_sleep"`
	Order         int           `json:"order"`
	ParsedObjects *[]PageObject `json:"parsed_objects"`
}

type PageObject struct {
	Key          string        `json:"key"`
	Selector     string        `json:"selector"`
	RegexExtract string        `json:"regex_extract"`
	PageObject   *[]PageObject `json:"objects"`
}

type Content interface {
}

type Source struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Data WebSite   `json:"data"`
}
