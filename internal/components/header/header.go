package header

import (
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/state"
)

func NewHeaderView(state *state.UIState) *tview.Flex {
	headerView := tview.NewFlex()
	headerView.SetBorder(true)

	tabView := NewTabView(state)
	statusView := NewStatusView(state)

	headerView.AddItem(tabView, 0, 1, true)
	headerView.AddItem(statusView, 0, 1, false)

	return headerView
}
