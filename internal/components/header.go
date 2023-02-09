package components

import (
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/state"
)

func NewHeaderView(state *state.UIState) *tview.Flex {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignRight).
			SetText(text)
	}

	headerView := tview.NewFlex()
	headerView.SetBorder(true)

	tabView := NewTabView(state)
	statusView := newPrimitive("Header")

	headerView.AddItem(tabView, 0, 1, true)
	headerView.AddItem(statusView, 0, 1, false)

	return headerView
}
