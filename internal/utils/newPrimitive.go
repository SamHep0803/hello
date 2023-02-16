package utils

import (
	"github.com/rivo/tview"
)

func NewPrimitive(text string) *tview.TextView {
	newPrimitive := tview.NewTextView()
	newPrimitive.SetTextAlign(tview.AlignCenter).
		SetBorder(true).
		SetTitle(text)

	return newPrimitive
}
