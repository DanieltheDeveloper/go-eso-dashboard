package component

import (
	"log"
	"net/http"

	"github.com/DanieltheDeveloper/go-eso-dashboard.git/pkg/constant"
	"github.com/PuerkitoBio/goquery"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// PlayerCountResponse is struct that represents the player count data.
type PlayerCountResponse struct {
	Current string `json:"current"`
	Peak    string `json:"24-peak"`
	AllPeak	string `json:"all-time-peak"`
}

const (
	Current = iota
	Peak
	AllPeak
)

// PeakPlayerCount is a component that displays the current player count.
type PeakPlayerCount struct {
	app.Compo
	PeakPlayerCount app.UI
}

// OnMount Check if the app is installable and set the state according.
func (p *PeakPlayerCount) OnMount(ctx app.Context) {
	p.PeakPlayerCount = app.Span().Text(getPeakPlayerCount(ctx))
}

// OnNav is called when the component is navigated to.
func (p *PeakPlayerCount) OnNav(ctx app.Context) {
	p.PeakPlayerCount = app.Span().Text(getPeakPlayerCount(ctx))
}

// Render is the main function that renders the current player count component.
func (p *PeakPlayerCount) Render() app.UI {
	if p.PeakPlayerCount == nil {
		return app.Span().Text("Loading...")
	}
	return p.PeakPlayerCount
}

// AllPeakPlayerCount is a component that displays the current player count.
type AllPeakPlayerCount struct {
	app.Compo
	AllPeakPlayerCount app.UI
}

// OnMount Check if the app is installable and set the state according.
func (a *AllPeakPlayerCount) OnMount(ctx app.Context) {
	a.AllPeakPlayerCount = app.Span().Text(getAllPeakPlayerCount(ctx))
}

// OnNav is called when the component is navigated to.
func (a *AllPeakPlayerCount) OnNav(ctx app.Context) {
	a.AllPeakPlayerCount = app.Span().Text(getAllPeakPlayerCount(ctx))
}

// Render is the main function that renders the current player count component.
func (a *AllPeakPlayerCount) Render() app.UI {
	if a.AllPeakPlayerCount == nil {
		return app.Span().Text("Loading...")
	}
	return a.AllPeakPlayerCount
}

// getPeakPlayerCount fetches the peak player count from Steam API.
func getPeakPlayerCount(ctx app.Context) (string) {
	// Check if state value is set and return
	peakPlayers := constant.Unreachable
	ctx.GetState("peakPlayersResponse", &peakPlayers)

	if peakPlayers != constant.Unreachable && peakPlayers != "" {
		return peakPlayers
	}

	peakPlayers = getSteamChartsPlayerCount(ctx).Peak
	ctx.SetState("peakPlayersResponse", peakPlayers).Persist().ExpiresIn(constant.PlayerCountCacheDuration) // TODO - ExpiresAt is always 0001-01-01T00:00:00Z inside local storage!?

	return peakPlayers
}

// getAllPeakPlayerCount fetches the peak player count from Steam API.
func getAllPeakPlayerCount(ctx app.Context) (string) {
	// Check if state value is set and return
	allPeakPlayers := constant.Unreachable
	ctx.GetState("allPeakPlayersResponse", &allPeakPlayers)

	if allPeakPlayers != constant.Unreachable && allPeakPlayers != "" {
		return allPeakPlayers
	}

	allPeakPlayers = getSteamChartsPlayerCount(ctx).AllPeak
	ctx.SetState("allPeakPlayersResponse", allPeakPlayers).Persist().ExpiresIn(constant.PlayerCountCacheDuration) // TODO - ExpiresAt is always 0001-01-01T00:00:00Z inside local storage!?

	return allPeakPlayers
}

// getSteamChartsPlayerCount fetches the player count from Steam Charts.
func getSteamChartsPlayerCount(ctx app.Context) (PlayerCountResponse) {
	// TODO - Add origins API to avoid CORS issues
	url := "https://api.allorigins.win/raw?url=https://steamcharts.com/app/306130"
	client := &http.Client{Timeout: constant.FetchTimeout}

	playerCountReturn := PlayerCountResponse{constant.Unreachable, constant.Unreachable, constant.Unreachable}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return playerCountReturn
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error creating request:", err)
		return playerCountReturn
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error fetching player count:", err)
		return playerCountReturn
	}

	doc.Find("#app-heading .num").Each(func(i int, s *goquery.Selection) {
		switch i {
		case Current:	// Current player count
			playerCountReturn.Current = s.Text()
		case Peak:	// 24h peak player count
			playerCountReturn.Peak = s.Text()
		case AllPeak:	// All time peak player count
			playerCountReturn.AllPeak = s.Text()
		default:
			return
		}
	})

	return playerCountReturn
}
