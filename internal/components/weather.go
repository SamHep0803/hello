package components

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

func NewWeatherComponent() *tview.TextView {
	weather := tview.NewTextView()

	weather.SetBorder(true).
		SetTitle("Weather")

	fmt.Fprintf(weather, "📌 Location: %s\n", viper.GetString("weather.city"))

	return weather
}
