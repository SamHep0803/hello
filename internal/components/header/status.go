package header

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/api"
	"github.com/samhep0803/hello/internal/creds"
	"github.com/samhep0803/hello/internal/state"
)

func NewStatusView(state *state.UIState) *tview.TextView {
	statusView := tview.NewTextView()

	statusView.SetTextAlign(tview.AlignRight)

	tokens, err := creds.GetCreds()
	githubToken := tokens["github"]
	if err != nil {
		return nil
	}

	githubName, err := api.GetGithubUsername(githubToken)
	if err != nil {
		return nil
	}

	fmt.Fprintf(statusView, "GitHub: %s", githubName)

	return statusView
}
