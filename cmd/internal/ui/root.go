package ui

import (
	"github.com/samhep0803/hello/cmd/internal/state"
)

func New() error {
	uiState := state.NewUIState()

	state.GlobalUIState = uiState

	// landingPage := pages.NewLandingPage(uiState)
	// uiState.Pages.AddPage(pages.LANDING_PAGE, landingPage, true, true)

	ui := NewUI(uiState)

	return uiState.App.SetRoot(ui, true).Run()
}
