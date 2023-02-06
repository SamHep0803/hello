package components

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/cmd/internal/state"
)

func NewTabView(state *state.UIState) *tview.TextView {
	tabView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			state.App.Draw()
		})
	tabView.Highlight()

	for index, tab := range state.Tabs {

	}

	tabView.SetDoneFunc(func(key tcell.Key) {
		currentSelection := tabView.GetHighlights()
		index, _ := strconv.Atoi(currentSelection[0])
		if key == tcell.KeyTAB {
			index = (index + 1) % len(currentSelection)
		} else if key == tcell.KeyBacktab {
			index = (index - 1 + len(currentSelection)) % len(currentSelection)
		}
		tabView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
	})

	return tabView
}
