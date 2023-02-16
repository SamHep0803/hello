package state

import (
	"github.com/rivo/tview"
)

var GlobalUIState *UIState

type Page struct {
	Title    string
	Contents tview.Primitive
}

type UIState struct {
	App        *tview.Application
	Pages      []Page
	CurrentTab int
	Content    *tview.Pages
}

func NewUIState() *UIState {
	return &UIState{
		App:        tview.NewApplication(),
		Pages:      []Page{},
		CurrentTab: 0,
		Content:    tview.NewPages(),
	}
}
