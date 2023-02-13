package header

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/state"
)

func NewTabView(state *state.UIState) *tview.TextView {
	tabView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			state.App.Draw()
		})
	tabView.Highlight(strconv.Itoa(state.CurrentTab))

	fmt.Fprintf(tabView, "[::b]TABS[-:-:-] |")
	for index, tab := range state.Tabs {
		fmt.Fprintf(tabView, ` ["%d"]%s[-][""] |`, index, tab.Title)
		state.Content.AddPage(strconv.Itoa(index), tab.Contents, true, index == state.CurrentTab)
	}

	tabView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTAB {
			state.CurrentTab = (state.CurrentTab + 1) % len(state.Tabs)
			state.Content.SwitchToPage(strconv.Itoa(state.CurrentTab))
		} else if key == tcell.KeyBacktab {
			state.CurrentTab = (state.CurrentTab - 1 + len(state.Tabs)) % len(state.Tabs)
			state.Content.SwitchToPage(strconv.Itoa(state.CurrentTab))
		}
		tabView.Highlight(strconv.Itoa(state.CurrentTab)).ScrollToHighlight()
	})

	tabView.SetHighlightedFunc(func(added, removed, remaining []string) {
		highlighted, _ := strconv.Atoi(added[0])
		state.CurrentTab = highlighted
		state.Content.SwitchToPage(added[0])
	})

	return tabView
}
