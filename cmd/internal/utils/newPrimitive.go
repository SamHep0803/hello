package utils

import (
	"fmt"

	"github.com/rivo/tview"
)

func NewPrimitive(text string) tview.Primitive {
	newPrimitive := tview.NewTextView()
	newPrimitive.SetTextAlign(tview.AlignCenter).
		SetBorder(true)

	fmt.Fprintf(newPrimitive, "%s", text)

	return newPrimitive
}
