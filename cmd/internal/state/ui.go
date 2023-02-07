package state

import (
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/cmd/internal/utils"
)

var GlobalUIState *UIState

type Tab struct {
	Title    string
	Contents tview.Primitive
}

type UIState struct {
	App        *tview.Application
	Tabs       []Tab
	CurrentTab int
	Content    *tview.Pages
}

func NewUIState() *UIState {
	return &UIState{
		App: tview.NewApplication(),
		Tabs: []Tab{
			{"Main", utils.NewPrimitive("Main Page")},
			{"Second", utils.NewPrimitive("Second Page")},
		},
		CurrentTab: 0,
		Content:    tview.NewPages(),
	}
}
