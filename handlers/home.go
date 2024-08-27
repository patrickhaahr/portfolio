package handlers

import (
	"GOTH/views/home"
	"net/http"
	"strconv"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("darkMode")
	if err == nil {
		isDarkMode, _ = strconv.ParseBool(cookie.Value)
	}
	return Render(w, r, home.Index(isDarkMode))
}

var isDarkMode bool // Track the toggle state globally

// ToggleIcon toggles between two icons and renders the correct one.
func ToggleIcon(w http.ResponseWriter, r *http.Request) error {
	isDarkMode = !isDarkMode // Toggle the state

	// Set a cookie to remember the user's preference
	http.SetCookie(w, &http.Cookie{
		Name:  "darkMode",
		Value: strconv.FormatBool(isDarkMode),
		Path:  "/",
	})

	// Send a custom header to trigger a full page reload
	w.Header().Set("HX-Refresh", "true")

	// Render the entire page
	return Render(w, r, home.Index(isDarkMode))
}
