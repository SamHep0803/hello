package components

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/rivo/tview"
)

type clockView struct {
	textView *tview.TextView
}

func (c *clockView) update(app *tview.Application) {
	for {
		now := time.Now()
		date := now.Format("Mon Jan 2 2006")

		app.Draw()
		c.textView.Clear()

		boldText := color.New(color.Bold)
		boldTime := boldText.Sprintf(now.Format("15:04:05"))

		c.textView.SetText(fmt.Sprintf("%s\n%s", tview.TranslateANSI(boldTime), date))

		app.Draw()

		time.Sleep(1 * time.Second)
		c.textView.Clear()
	}
}

func NewClockView(app *tview.Application) *tview.TextView {
	textView := tview.NewTextView()
	textView.SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	view := &clockView{
		textView: textView,
	}

	go view.update(app)

	return view.textView
}
