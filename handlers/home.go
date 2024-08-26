package handlers

import (
	"GOTH/theme"
	"GOTH/views/home"
	"net/http"

	"github.com/a-h/templ"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}

var isDarkMode bool // Track the toggle state globally

// ToggleIcon toggles between two icons and renders the correct one.
func ToggleIcon(w http.ResponseWriter, r *http.Request) error {
	isDarkMode = !isDarkMode // Toggle the state
	var icon templ.Component
	if isDarkMode {
		icon = theme.DarkModeIcon()
	} else {
		icon = theme.LightModeIcon()
	}
	return icon.Render(r.Context(), w)
}
