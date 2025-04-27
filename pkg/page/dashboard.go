package page

// TODO - Convert javascript to async go routines
// TODO - Get Server Status from API, set as local storage and use it to update the UI every 3 minutes (Without site refresh?! https://go-app.dev/components)

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// dashboard is a component that displays a simple ESO Dashboard. A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Dashboard struct {
	app.Compo

	isAppInstallable bool

	rssFeed      app.UI
	serverStatus []string
	activeUsers  int
	peak24h      int
	allPeak      int
}

// RssFeed is struct that represents the RSS feed data.
type RSSFeedResponse struct {
	Items []struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Thumbnail   string `json:"thumbnail"`
		PubDate     string `json:"pubDate"`
	} `json:"items"`
}

// Check if the app is installable and set the state accordingly
func (d *Dashboard) OnMount(ctx app.Context) {
	d.isAppInstallable = ctx.IsAppInstallable()
}

// Check if the app is installable and set the state accordingly
func (d *Dashboard) OnAppInstallChange(ctx app.Context) {
	d.isAppInstallable = ctx.IsAppInstallable()
}

// Reload data every time the page is loaded or refreshed
func (d *Dashboard) OnPreRender(ctx app.Context) {
	ctx.Async(func() {
		d.fetchRSSFeed(ctx)
	})
}

// Reload data every time the page is loaded or refreshed
func (d *Dashboard) OnNav(ctx app.Context) {
	ctx.Async(func() {
		d.fetchRSSFeed(ctx)
	})
}

// OnAppUpdate satisfies the app.AppUpdater interface. It is called when the app
// is updated in background.
func (d *Dashboard) OnAppUpdate(ctx app.Context) {
	ctx.Reload() // Reload the app when it is updated
}

// The Render method is where the component appearance is defined.
func (d *Dashboard) Render() app.UI {
	return app.Div().ID("dashboard").Body(
		app.Video().Style("width", "110vw").Style("height", "110vh").Style("object-fit", "cover").Style("position", "fixed").Style("z-index", "-1").Style("top", "0").Style("left", "0").
			ID("bg-video").Muted(true).Loop(true).AutoPlay(true).Src("/web/background-video.mp4"),
		app.Div().Class("container mt-6").Body(
			app.H1().Class("text-center p-2").Text("Elder Scrolls Online"),
			app.Div().Style("width", "100%").Style("height", "100%").Class("row mt-4 d-flex align-items-stretch").Body(
				// System Status Card
				app.Div().Class("col-md-4 mt-4").Body(
					app.Div().Class("card flex d-flex flex-column h-100").Body(
						app.Div().Class("card-header text-center bg-primary text-white").Text("ESO Server Status"),
						app.Div().Class("card-body flex-grow-1 text-center flex-row d-flex align-items-center justify-content-center m-2").Body(
							app.Div().Style("width", "100%").Body(
								app.Ul().Class("list-group").ID("serverStatusList").Body(
									app.Li().Class("list-group-item text-white bg-dark").Text("Loading..."),
								),
								app.Br(),
								app.P().Class("card-text").ID("status").Text("Loading..."),
							),
						),
					),
				),
				// RSS Feed Card
				app.Div().Class("col-md-8 mt-4").Body(
					app.Div().Class("card d-flex flex-column h-100").Body(
						app.Div().Class("card-header text-center bg-success text-white").Text("Latest ESO News"),
						app.Div().ID("rss-feed").
							Class("card-body flex-grow-1 m-4").
							Body(
								d.rssFeed,
							),
					),
				),
			),
			// Additional Stats
			app.Div().Style("width", "100%").Class("row mt-4").Body(
				app.Div().Class("col-md-4").Body(
					app.Div().Class("card text-center").Body(
						app.Div().Class("card-header bg-info text-white").Text("Active Users"),
						app.Div().Class("card-body").Body(
							app.H5().Class("card-title").ID("activeUsers").Text("0"),
						),
					),
				),
				app.Div().Class("col-md-4").Body(
					app.Div().Class("card text-center").Body(
						app.Div().Class("card-header bg-warning text-white").Text("24 Hour Peak"),
						app.Div().Class("card-body").Body(
							app.H5().Class("card-title").ID("24Peak").Text("0"),
						),
					),
				),
				app.Div().Class("col-md-4").Body(
					app.Div().Class("card text-center").Body(
						app.Div().Class("card-header bg-danger text-white").Text("All-Time Peak"),
						app.Div().Class("card-body").Body(
							app.H5().Class("card-title").ID("allPeak").Text("0"),
						),
					),
				),
			),
		),
	)
}

// Cache RSS feed data in state, so it doesn't need to be fetched every time the page is loaded
func (d *Dashboard) fetchRSSFeed(ctx app.Context) {
	// Check if state value is set and return
	rssFeedResponse := RSSFeedResponse{}
	ctx.GetState("rssFeedResponse", &rssFeedResponse)
	if len(rssFeedResponse.Items) == 0 {
		// Make API request
		url := "https://api.rss2json.com/v1/api.json?rss_url=https://eso-hub.com/en/news/feed.rss"
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error making request:", err)
			d.rssFeed = app.Span().Text("Error fetching RSS feed")
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response:", err)
			d.rssFeed = app.Span().Text("Error fetching RSS feed")
			return
		}

		err = json.Unmarshal(body, &rssFeedResponse)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			d.rssFeed = app.Span().Text("Error fetching RSS feed")
			return
		}
	}

	// Set state value to expire in 24 hours
	ctx.SetState("rssFeedResponse", rssFeedResponse).Persist().ExpiresAt(time.Now().Add(24 * time.Hour))

	// Wrap the items in a div
	div := app.Div().Class("rss-feed").Class("d-flex flex-column gap-3")
	itemsDiv := make([]app.UI, 3) // Limit to 3 items
	for i := 0; i < len(rssFeedResponse.Items) && i < 3; i++ {
		item := rssFeedResponse.Items[i]
		// Parse the date string
		pubDate, err := time.Parse("2006-01-02 15:04:05", item.PubDate)
		if err != nil {
			log.Println("Error parsing date: ", err)
		}

		itemsDiv[i] = app.Div().Class("relative block bg-gray-900 border border-gray-900 shadow-lg rounded overflow-hidden").
			Body(
				app.A().Href(item.Link).Target("_blank").Style("text-decoration", "none").
					Body(
						app.Img().Style("width", "100%").Src(item.Thumbnail).Alt(item.Title),
						app.Div().Class("absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black via-black/60 to-transparent").
							Body(
								app.H3().Class("text-white text-2xl font-semibold shadow-black line-clamp-2").Text(item.Title).Style("text-shadow", "black 2px 2px 1px"),
								app.P().Class("text-white mt-2 line-clamp-2 text-sm").Text(item.Description),
								app.P().Class("text-white text-sm mt-2 line-clamp-1").Text("Published on "+pubDate.Format("2006-01-02")),
							),
					),
			)
	}

	div.Body(itemsDiv...)
	d.rssFeed = div
}
