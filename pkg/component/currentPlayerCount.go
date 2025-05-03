package component

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DanieltheDeveloper/go-eso-dashboard.git/pkg/constant"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// CurrentPlayers is a component that displays the current player count.
type CurrentPlayers struct {
	app.Compo
	CurrentPlayers app.UI
}

// OnMount Check if the app is installable and set the state according.
func (c *CurrentPlayers) OnMount(ctx app.Context) {
	c.CurrentPlayers = app.Span().Text(getCurrentPlayers(ctx))
}

// OnNav is called when the component is navigated to.
func (c *CurrentPlayers) OnNav(ctx app.Context) {
	c.CurrentPlayers = app.Span().Text(getCurrentPlayers(ctx))
}

// Render is the main function that renders the current player count component.
func (c *CurrentPlayers) Render() app.UI {
	if c.CurrentPlayers == nil {
		return app.Span().Text("Loading...")
	}
	return c.CurrentPlayers
}

// getCurrentPlayers fetches the current player count from Steam API.
func getCurrentPlayers(ctx app.Context) (string) {
	// Check if state value is set and return
	currentPlayers := constant.Unreachable
	ctx.GetState("currentPlayersResponse", &currentPlayers)

	if currentPlayers != constant.Unreachable && currentPlayers != "" {
		return currentPlayers
	}

	// TODO - Add origins API to avoid CORS issues
	url := "https://api.allorigins.win/raw?url=" +
		"https://api.steampowered.com/ISteamUserStats/GetNumberOfCurrentPlayers/v1/?appid=306130"
	client := &http.Client{Timeout: constant.FetchTimeout}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return constant.Unreachable
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error fetching current player count:", err)
		return constant.Unreachable
	}
	defer resp.Body.Close()

	var data struct {
		Response struct {
			PlayerCount int `json:"player_count"`
		} `json:"response"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		// TODO - Sometimes the return is broken with error <
		log.Println("Error decoding response:", err)
		return constant.Unreachable
	}

	ctx.SetState("currentPlayersResponse", strconv.Itoa(data.Response.PlayerCount)).
	Persist().ExpiresIn(constant.PlayerCountCacheDuration) // TODO - ExpiresAt is always 0001-01-01T00:00:00Z inside local storage!?

	return strconv.Itoa(data.Response.PlayerCount)
}