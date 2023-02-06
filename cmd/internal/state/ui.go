package state

import "github.com/rivo/tview"

var GlobalUIState *UIState

type Tab struct {
	Title    string
	Contents tview.Primitive
}

type UIState struct {
	App        *tview.Application
	Tabs       []Tab
	CurrentTab int
}

func NewUIState() *UIState {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	return &UIState{
		App: tview.NewApplication(),
		Tabs: []Tab{
			{"Main", newPrimitive("Main Page")},
			{"Second", newPrimitive("Second Page")},
		},
		CurrentTab: 0,
	}
}
