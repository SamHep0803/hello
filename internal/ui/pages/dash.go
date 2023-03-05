package pages

import (
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/components"
	"github.com/samhep0803/hello/internal/state"
)

func NewDashboardPage(state *state.UIState) tview.Primitive {
	dash := tview.NewFlex()

	dash.SetBorder(true).SetTitle("Dash Flex")

	// weather view
	weatherView := components.NewWeatherComponent()

	dash.AddItem(weatherView, 40, 1, false)

	return dash
}
