package state

import "github.com/rivo/tview"

var GlobalUIState *UIState

type UIState struct {
	App   *tview.Application
	Pages *tview.Pages
}

func NewUIState() *UIState {
	return &UIState{
		App:   tview.NewApplication(),
		Pages: tview.NewPages(),
	}
}
