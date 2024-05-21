package domain

type IPage interface {
	GetPageEvents() []PageObject
	SetPage()
}
