package ui

import (
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/cmd/internal/components"
	"github.com/samhep0803/hello/cmd/internal/state"
	"github.com/samhep0803/hello/cmd/internal/utils"
)

func NewUI(state *state.UIState) tview.Primitive {
	clockView := components.NewClockView(state.App)
	headerView := components.NewHeaderView(state)
	mainView := state.Content

	menu := utils.NewPrimitive("Menu")
	sideBar := utils.NewPrimitive("Side Bar")

	page := tview.NewGrid().
		SetRows(4, 0, 3).
		SetColumns(20, 0, 30).
		// SetBorders(true).
		AddItem(utils.NewPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	page.AddItem(mainView, 1, 0, 1, 3, 0, 0, false).
		AddItem(headerView, 0, 0, 1, 3, 0, 0, true).
		AddItem(clockView, 0, 0, 1, 0, 0, 0, false)

	// Layout < 50
	page.AddItem(clockView, 0, 0, 1, 1, 0, 75, false).
		AddItem(headerView, 0, 1, 1, 2, 0, 75, true)

	// Layout for screens wider than 100 cells.
	page.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(mainView, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false).
		AddItem(clockView, 0, 0, 1, 1, 0, 100, false).
		AddItem(headerView, 0, 1, 1, 2, 0, 100, true)

	return page
}
