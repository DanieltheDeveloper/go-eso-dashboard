package rssFeed

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// RSSFeed is a component that displays an RSS feed.
type RSSFeed struct {
	app.Compo
	RSSFeed app.UI
}

// maxRSSItems is the maximum number of RSS items to display.
const maxRSSItems = 3

// rssFeedCacheDuration is the duration for which RSS feed data is cached.
const rssFeedCacheDuration = 24 * time.Hour

// RSSFeedResponse is struct that represents the RSS feed data.
type rssFeedResponse struct {
	Items []struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Thumbnail   string `json:"thumbnail"`
		PubDate     string `json:"pubDate"`
	} `json:"items"`
}

// OnMount Check if the app is installable and set the state according.
func (r *RSSFeed) OnMount(ctx app.Context) {
	r.RSSFeed = fetchRSSFeed(ctx)
}

// OnNav is called when the component is navigated to.
func (r *RSSFeed) OnNav(ctx app.Context) {
    r.RSSFeed = fetchRSSFeed(ctx)
}

// Render is the main function that renders the RSS feed component.
func (r *RSSFeed) Render() app.UI {
	return app.Div().Body(r.RSSFeed)
}

// fetchRSSFeed Cache RSS feed data in state, so it doesn't need to be fetched every time the page is loaded.
func fetchRSSFeed(ctx app.Context) app.UI {
	// Check if state value is set and return
	rssFeed := rssFeedResponse{}
	ctx.GetState("rssFeedResponse", &rssFeed)
	if len(rssFeed.Items) == 0 {
		// Make API request
		url := "https://api.rss2json.com/v1/api.json?rss_url=https://eso-hub.com/en/news/feed.rss"
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			log.Println("Error creating request:", err)
			return app.Span().Text("Error fetching RSS feed")
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Error making request:", err)
			return app.Span().Text("Error fetching RSS feed")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response:", err)
			return app.Span().Text("Error fetching RSS feed")
		}

		err = json.Unmarshal(body, &rssFeed)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			return app.Span().Text("Error fetching RSS feed")
		}
	}

	// Set state value to expire in 24 hours
	ctx.SetState("rssFeedResponse", rssFeed).Persist().ExpiresAt(time.Now().Add(rssFeedCacheDuration))

	// Wrap the items in a div
	div := app.Div().Class("rss-feed").Class("d-flex flex-column gap-3")
	itemsDiv := make([]app.UI, maxRSSItems) // Limit to 3 items
	for i := 0; i < len(rssFeed.Items) && i < 3; i++ {
		item := rssFeed.Items[i]
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
								app.H3().Class("text-white text-2xl font-semibold shadow-black line-clamp-2").
								Text(item.Title).Style("text-shadow", "black 2px 2px 1px"),
								app.P().Class("text-white mt-2 line-clamp-2 text-sm").Text(item.Description),
								app.P().Class("text-white text-sm mt-2 line-clamp-1").Text("Published on "+pubDate.Format("2006-01-02")),
							),
					),
			)
	}

	div.Body(itemsDiv...)
	return div
}
