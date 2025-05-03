package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DanieltheDeveloper/go-eso-dashboard.git/pkg/page"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const (
	serverReadWriteTimeout	= 15
	serverIdleTimeout       = 60
	serverReadHeaderTimeout = 10
)

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the dashboard component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", func() app.Composer { return &page.Dashboard{} })

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:         "ESO Dashboard",
		Title:        "ESO Dashboard",
		ShortName:    "ESO Dashboard",
		LoadingLabel: "ESO Dashboard data is loading ... {progress}%",
		Lang:         "en",
		Author:       "DanielTheDeveloper",
		Description:  "Simple Go ESO dashboard with caching support for local deployment",
		Icon: app.Icon{
			Default:  "web/eso.png",
			Large:    "web/eso.png",
			SVG:      "web/eso.svg",
			Maskable: "web/eso-maskable.png",
		},
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css",
			"/web/eso-dashboard.css",
		},
		CacheableResources: []string{
			"/web/background-video.mp4",
		},
	})
	
	// Create a server with proper timeout settings
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           nil,
		ReadTimeout:       serverReadWriteTimeout * time.Second,
		WriteTimeout:      serverReadWriteTimeout * time.Second,
		IdleTimeout:       serverIdleTimeout * time.Second,
		ReadHeaderTimeout: serverReadHeaderTimeout * time.Second,
	}

	// Start the server with proper timeout configurations
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
