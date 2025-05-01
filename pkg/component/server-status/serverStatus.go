package serverStatus

import (
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// serverStatusResponse is struct that represents the server status data.
type serverStatusResponse struct {
	PCEU    string `json:"PC-EU"`
	PCNA    string `json:"PC-NA"`
	PCPTS 	string `json:"PC-PTS"`
	XBOXEU  string `json:"XBOX-EU"`
	XBOXNA  string `json:"XBOX-NA"`
	PS4NA	string `json:"PS4-NA"`
	PS4EU	string `json:"PS4-EU"`
}

// ServerStatusType represents the different states a server can be in.
type serverStatusType string

const (
	ServerStatusOperational      serverStatusType = "All servers operational"
	ServerStatusMinorIssues      serverStatusType = "Minor issues detected"
	ServerStatusMaintenance      serverStatusType = "Server maintenance ongoing"
	ServerStatusCritical         serverStatusType = "Critical failure detected"
)

func (s serverStatusType) String() string {
	return string(s)
}

// ServerRegion represents the different regions a server can be in.
type serverRegion string

const (
	PCEU   		serverRegion = "PC-EU"
	PCNA  		serverRegion = "PC-NA"
	PCPTS       serverRegion = "PC-PTS"
	XBOXEU  	serverRegion = "XBOX-EU"
	XBOXNA 		serverRegion = "XBOX-NA"
	PS4NA		serverRegion = "PS4-NA"
	PS4EU   	serverRegion = "PS4-EU"
)

func (s serverRegion) String() string {
	return string(s)
}

// RSSFeed is a component that displays an RSS feed.
type ServerStatus struct {
	app.Compo
	ServerStatus app.UI
}

// serverStatusCacheDuration is the duration for which the data is cached.
const serverStatusCacheDuration = 5 * time.Minute

// serverStatusCacheDuration is the duration for which the data is cached.
const fetchTimeout = 10 * time.Second

// OnMount Check if the app is installable and set the state according.
func (s *ServerStatus) OnMount(ctx app.Context) {
	s.ServerStatus = fetchServerStatus(ctx)
}

// OnNav is called when the component is navigated to.
func (s *ServerStatus) OnNav(ctx app.Context) {
	s.ServerStatus = fetchServerStatus(ctx)
}

// Render is the main function that renders the ServerStatus component.
func (s *ServerStatus) Render() app.UI {
	return app.Div().Body(s.ServerStatus)
}

// getStatusClass returns the class name based on the server status.
func getStatusClass(status string) string {
	switch status {
	case "Online":
		return "list-group-item bg-success"
	case "Offline":
		return "list-group-item bg-warning"
	default:
		return "list-group-item bg-danger"
	}
}

// fetchServerStatus Cache data in state, so it doesn't need to be fetched every time the page is loaded.
func fetchServerStatus(ctx app.Context) app.UI {
	// Check if state value is set and return
	serverStatus := serverStatusResponse{}
	ctx.GetState("serverStatusResponse", &serverStatus)

	// Create a list to hold the server status items
	statusList := make([]app.UI, reflect.TypeOf(serverStatusResponse{}).NumField())

	// Check if any entries in struct are empty
	if serverStatus.PCEU != "" && serverStatus.PCNA != "" && serverStatus.PCPTS != "" &&
	serverStatus.XBOXEU != "" && serverStatus.XBOXNA != "" && 
	serverStatus.PS4NA != "" && serverStatus.PS4EU != "" {
		// Return cached response
		// Loop through each struct member and retrieve the cached status
		statusList = []app.UI{
			app.Li().Class(getStatusClass(serverStatus.PCEU)).
			Body(app.Text(string(PCEU) + ": " + serverStatus.PCEU)),
			app.Li().Class(getStatusClass(serverStatus.PCNA)).
			Body(app.Text(string(PCNA) + ": " + serverStatus.PCNA)),
			app.Li().Class(getStatusClass(serverStatus.PCPTS)).
			Body(app.Text(string(PCPTS) + ": " + serverStatus.PCPTS)),
			app.Li().Class(getStatusClass(serverStatus.XBOXEU)).
			Body(app.Text(string(XBOXEU) + ": " + serverStatus.XBOXEU)),
			app.Li().Class(getStatusClass(serverStatus.XBOXNA)).
			Body(app.Text(string(XBOXNA) + ": " + serverStatus.XBOXNA)),
			app.Li().Class(getStatusClass(serverStatus.PS4NA)).
			Body(app.Text(string(PS4NA) + ": " + serverStatus.PS4NA)),
			app.Li().Class(getStatusClass(serverStatus.PS4EU)).
			Body(app.Text(string(PS4EU) + ": " + serverStatus.PS4EU)),
		}
		return app.Div().Body(statusList...)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: fetchTimeout,
	}

	// Make request through CORS proxy
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.allorigins.win/raw?url=https://esoserverstatus.net/",
		nil,
		)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return app.Div().Text("Error creating request")
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching server status: %v", err)
		return app.Div().Text("Error loading server status")
	}
	defer resp.Body.Close()

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing response: %v", err)
		return app.Div().Text("Error parsing server status")
	}

	// Check each server region
	for key, region := range []serverRegion{PCEU, PCNA, PCPTS, XBOXEU, XBOXNA, PS4NA, PS4EU} {
		selector := "#" + string(region)
		status := "Unknown"
		
		doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
			if b := s.Find("b"); b.Length() > 0 {
				status = b.Text()
			}
		})

		// Set the status in the struct based on the current region
		switch region {
		case PCEU:
			serverStatus.PCEU = status
		case PCNA:
			serverStatus.PCNA = status
		case PCPTS:
			serverStatus.PCPTS = status
		case XBOXEU:
			serverStatus.XBOXEU = status
		case XBOXNA:
			serverStatus.XBOXNA = status
		case PS4NA:
			serverStatus.PS4NA = status
		case PS4EU:
			serverStatus.PS4EU = status
		}

		statusList[key] = app.Li().Class(getStatusClass(status)).Body(app.Text(string(region) + ": " + status))
	}
	
	// Check if statusList is empty
	if len(statusList) == 0 {
		return app.Div().Text("Error parsing server status")
	}

	statusDiv := app.Div().Body(statusList...)

	// Cache the result
	ctx.SetState("serverStatusResponse", serverStatus).Persist().ExpiresIn(serverStatusCacheDuration)

	return statusDiv
}
