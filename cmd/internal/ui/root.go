package ui

import (
	"github.com/samhep0803/hello/cmd/internal/state"
	"github.com/samhep0803/hello/cmd/internal/ui/pages"
)

func New() error {
	uiState := state.NewUIState()

	state.GlobalUIState = uiState

	landingPage := pages.NewLandingPage(uiState)
	uiState.Pages.AddPage(pages.LANDING_PAGE, landingPage, true, true)

	return uiState.App.SetRoot(uiState.Pages, true).Run()
}
