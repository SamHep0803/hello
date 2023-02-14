package header

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/api"
	"github.com/samhep0803/hello/internal/creds"
	"github.com/samhep0803/hello/internal/state"
	"github.com/spf13/viper"
)

const refreshInterval = 15 * time.Minute

func NewStatusView(state *state.UIState) *tview.TextView {
	statusView := tview.NewTextView()

	statusView.SetTextAlign(tview.AlignRight)

	tokens, err := creds.GetCreds()
	githubToken := tokens["github"]
	weatherToken := tokens["weather"]
	if err != nil {
		return nil
	}

	githubName, err := api.GetGithubUsername(githubToken)
	if err != nil {
		return nil
	}

	coords, err := api.GetCoordinatesByCity(viper.GetString("weather.city"), weatherToken)
	if err != nil {
		return nil
	}
	lat := coords["lat"]
	lon := coords["lon"]

	temp, _ := api.GetTemperature(lon, lat, weatherToken)

	go func() {
		for {
			time.Sleep(refreshInterval)
			t, _ := api.GetTemperature(lon, lat, weatherToken)
			temp = t
		}
	}()

	fmt.Fprintf(statusView, "GitHub: %s\n", githubName)
	fmt.Fprintf(statusView, "Weather: %s, %s°C\n", viper.GetString("weather.city"), fmt.Sprint(temp))

	return statusView
}
