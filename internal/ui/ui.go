package ui

import (
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/components"
	"github.com/samhep0803/hello/internal/components/header"
	"github.com/samhep0803/hello/internal/state"
	"github.com/samhep0803/hello/internal/utils"
)

func NewUI(state *state.UIState) tview.Primitive {
	clockView := components.NewClockView(state.App)
	headerView := header.NewHeaderView(state)
	mainView := state.Content

	ui := tview.NewGrid().
		SetRows(4, 0, 3).
		SetColumns(20, 0, 30).
		// SetBorders(true).
		AddItem(utils.NewPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	ui.AddItem(mainView, 1, 0, 1, 3, 0, 0, false).
		AddItem(headerView, 0, 0, 1, 3, 0, 0, true).
		AddItem(clockView, 0, 0, 1, 0, 0, 0, false)

	// Layout < 50
	ui.AddItem(clockView, 0, 0, 1, 1, 0, 75, false).
		AddItem(headerView, 0, 1, 1, 2, 0, 75, true)

	// Layout for screens wider than 100 cells.
	ui.AddItem(mainView, 1, 0, 1, 3, 0, 100, false).
		AddItem(clockView, 0, 0, 1, 1, 0, 100, false).
		AddItem(headerView, 0, 1, 1, 2, 0, 100, true)

	return ui
}
