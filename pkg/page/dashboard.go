package page

// TODO - Convert javascript to async go routines
// TODO - Get Server Status from API, set as local storage and use it to update the UI every 3 minutes (Without site refresh?! https://go-app.dev/components)

import (
	"github.com/DanieltheDeveloper/go-eso-dashboard.git/pkg/component"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Dashboard is a component that displays a simple ESO Dashboard. A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Dashboard struct {
	app.Compo
	RSSFeed component.RSSFeed
	ServerStatus component.ServerStatus
	CurrentPlayers component.CurrentPlayers
	PeakPlayerCount component.PeakPlayerCount
	AllPeakPlayerCount component.AllPeakPlayerCount
	isAppInstallable bool
}

// OnMount Check if the app is installable and set the state according.
func (d *Dashboard) OnMount(ctx app.Context) {
	d.isAppInstallable = ctx.IsAppInstallable()
}

// OnAppInstallChange Check if the app is installable and set the state accordingly.
func (d *Dashboard) OnAppInstallChange(ctx app.Context) {
	d.isAppInstallable = ctx.IsAppInstallable()
}

// OnAppUpdate is called when the app is updated in background.
func (d *Dashboard) OnAppUpdate(ctx app.Context) {
	ctx.Reload() // Reload the app when it is updated
}

// Render The Render method is where the component appearance is defined.
func (d *Dashboard) Render() app.UI {
	return app.Div().ID("dashboard").Body(
		app.Video().Style("width", "110vw").Style("height", "110vh").Style("object-fit", "cover").
		Style("position", "fixed").Style("z-index", "-1").Style("top", "0").Style("left", "0").
			ID("bg-video").Muted(true).Loop(true).AutoPlay(true).Src("/web/background-video.mp4"),
		app.Div().Class("container mt-6").Body(
			app.H1().Class("text-center p-2").Text("Elder Scrolls Online"),
			app.Div().Style("width", "100%").Style("height", "100%").Class("row mt-4 d-flex align-items-stretch").Body(
				// System Status Card
				app.Div().Class("col-md-4 mt-4").Body(
					app.Div().Class("card flex d-flex flex-column h-100").Body(
						app.Div().Class("card-header text-center bg-primary text-white").Text("ESO Server Status"),
						app.Div().Class(
							"card-body flex-grow-1 text-center flex-row d-flex align-items-center justify-content-center m-2",
							).Body(
							app.Div().Style("width", "100%").Body(
								app.Ul().Class("list-group").ID("serverStatusList").Body(
									&d.ServerStatus,
								),
							),
						),
					),
				),
				// RSS Feed Card
				app.Div().Class("col-md-8 mt-4").Body(
					app.Div().Class("card d-flex flex-column h-100").Body(
						app.Div().Class("card-header text-center bg-success text-white").Text("Latest ESO News"),
						app.Div().ID("rss-feed").
							Class("card-body flex-grow-1 m-4").Body(&d.RSSFeed),
					),
				),
			),
			// Additional Stats
			app.Div().Style("width", "100%").Class("row mt-4").Body(
				app.Div().Class("col-md-4").Body(
					app.Div().Class("card text-center").Body(
						app.Div().Class("card-header bg-info text-white").Text("Active Users"),
						app.Div().Class("card-body").Body(
							app.H5().Class("card-title").ID("activeUsers").Body(&d.CurrentPlayers),
						),
					),
				),
				app.Div().Class("col-md-4").Body(
					app.Div().Class("card text-center").Body(
						app.Div().Class("card-header bg-warning text-white").Text("24 Hour Peak"),
						app.Div().Class("card-body").Body(
							app.H5().Class("card-title").ID("24Peak").Body(&d.PeakPlayerCount),
						),
					),
				),
				app.Div().Class("col-md-4").Body(
					app.Div().Class("card text-center").Body(
						app.Div().Class("card-header bg-danger text-white").Text("All-Time Peak"),
						app.Div().Class("card-body").Body(
							app.H5().Class("card-title").ID("allPeak").Body(&d.AllPeakPlayerCount),
						),
					),
				),
			),
		),
	)
}