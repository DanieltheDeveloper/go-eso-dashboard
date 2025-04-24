package page

// TODO - Convert javascript to async go routines to allow ui Goroutine to update

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// dashboard is a component that displays a simple ESO Dashboard. A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Dashboard struct {
	app.Compo

	serverStatus []string
	rssFeed      []string
	activeUsers  int
	peak24h      int
	allPeak      int
}

func (d *Dashboard) OnPreRender(ctx app.Context) {
	// TODO - Fetch server status and RSS feed data ... from the API
}

func (d *Dashboard) OnNav(ctx app.Context) {
	// TODO - Fetch server status and RSS feed data ... from the API
}

// The Render method is where the component appearance is defined.
func (d *Dashboard) Render() app.UI {
	return app.Div().ID("dashboard").Body(
		app.Video().Style("width", "110vw").Style("height", "110vh").Style("object-fit", "cover").Style("position", "fixed").Style("z-index", "-1").Style("top", "0").Style("left", "0").
			ID("bg-video").Muted(true).Loop(true).AutoPlay(true).Controls(false).Src("/web/background-video.mp4"),
		app.Div().Class("container mt-6").Body(
			app.H1().Class("text-center p-2").Text("Elder Scrolls Online"),
			app.Div().Style("width", "100%").Style("height", "100%").Class("row mt-4 d-flex align-items-stretch").Body(
				// System Status Card
				app.Div().Class("col-md-4").Body(
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
				app.Div().Class("col-md-8").Body(
					app.Div().Class("card d-flex flex-column h-100").Body(
						app.Div().Class("card-header text-center bg-success text-white").Text("Latest ESO News"),
						app.Div().Class("card-body flex-grow-1 m-2").Style("max-height", "620px").Style("overflow-y", "auto").Body(
							app.Div().ID("rssFeed").Class("d-flex flex-column gap-3").Body(
								app.Div().Class("text-white bg-dark").Text("Loading..."),
							),
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
