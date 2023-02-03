package pages

import (
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/cmd/internal/components"
	"github.com/samhep0803/hello/cmd/internal/state"
)

var LANDING_PAGE = "landing_page"

func NewLandingPage(state *state.UIState) tview.Primitive {
	clockView := components.NewClockView(state.App)

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")
	sideBar := newPrimitive("Side Bar")

	page := tview.NewGrid().
		SetRows(2, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	header := newPrimitive("Header")

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	page.AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(header, 0, 0, 1, 3, 0, 0, false)

	// Layout for screens wider than 100 cells.
	page.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false).
		AddItem(clockView, 0, 0, 1, 1, 0, 100, false).
		AddItem(header, 0, 1, 1, 2, 0, 100, false)

	return page
}
