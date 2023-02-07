package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/samhep0803/hello/cmd/internal/state"
)

func New() error {
	uiState := state.NewUIState()

	state.GlobalUIState = uiState

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
