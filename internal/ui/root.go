package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/samhep0803/hello/internal/state"
	"github.com/samhep0803/hello/internal/ui/pages"
	"github.com/samhep0803/hello/internal/utils"
)

func New() error {
	uiState := state.NewUIState()

	state.GlobalUIState = uiState

	uiState.Pages = []state.Page{
		{Title: "Dashboard", Contents: pages.NewDashboardPage(uiState)},
		{Title: "Second", Contents: utils.NewPrimitive("Second Page")},
	}

	ui := NewUI(uiState)

	uiState.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			uiState.App.Stop()
		}
		return event
	})

	uiState.App.EnableMouse(true)

	return uiState.App.SetRoot(ui, true).Run()
}
