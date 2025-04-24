package page

// TODO - Put css in raw headers or new .css file?
// TODO - Add javascript to js file and load with app header

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// dashboard is a component that displays a simple ESO Dashboard. A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Dashboard struct {
	app.Compo
}

// The Render method is where the component appearance is defined.
func (d *Dashboard) Render() app.UI {
	return app.Div().Class("container mt-6").Body(
		// Include Bootstrap CSS
		app.Raw(`
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet">
		`),
		// Embed CSS styling
		app.Raw(`
			<style>
				@import url('https://fonts.googleapis.com/css2?family=Trajan+Pro:wght@400;700&display=swap');
				@import url('https://fonts.googleapis.com/css2?family=Cinzel:wght@400;700&display=swap');

				body {
				display: flex;
				justify-content: center;
				align-items: center;
				min-height: 100vh;
				margin: 0;
				background-color: #000;
				color: #c7b377;
				font-family: 'Cinzel', serif;
				background-image: url('https://images.ctfassets.net/rporu91m20dc/4gNjttrkh46UWUQCkkOAsm/eeff79c08bca3d6633dd73f37b03f413/ESO_Necrom_Key_Art_Carousel_Desktop.jpg');
				background-size: cover;
				background-attachment: fixed;
				background-position: center;
			}

			.container {
				margin: 2rem; /* Remove auto margin */
				width: 90%;
				max-width: 1200px;
				padding: 2rem;
				background-color: rgba(0, 0, 0, 0.9);
				border: 1px solid #8e7c44;
				box-shadow: 0 0 20px rgba(199, 179, 119, 0.1);
			}

				/* Mobile responsiveness */
				@media (max-width: 768px) {
					.container {
						width: 95%;
						padding: 1rem;
						position: relative;
						transform: none;
						top: 0;
						left: 0;
						margin: 1rem auto;
					}

					h1 {
						font-size: 1.8rem;
						margin-bottom: 1rem;
					}

					.card {
						margin-bottom: 0.5rem;
					}

					.card-header {
						padding: 0.5rem;
					}

					canvas {
						height: 200px !important;
						width: 100% !important;
					}
				}

				/* Small mobile devices */
				@media (max-width: 576px) {
					h1 {
						font-size: 1.5rem;
					}

					.container {
						padding: 0.5rem;
					}

					.row {
						margin-left: -5px;
						margin-right: -5px;
					}

					.col-md-4,
					.col-md-8,
					.col-md-6 {
						padding: 5px;
					}
				}

				.card {
					background-color: rgba(20, 20, 20, 0.95);
					border: 1px solid #8e7c44;
					margin-bottom: 1rem;
					box-shadow: 0 0 10px rgba(199, 179, 119, 0.2);
				}

				.card-header {
					background-color: rgba(142, 124, 68, 0.3) !important;
					color: #c7b377 !important;
					border-bottom: 1px solid #8e7c44;
					font-weight: bold;
					text-transform: uppercase;
					letter-spacing: 1px;
				}

				.card-body {
					color: #c7b377;
				}

				h1 {
				color: #c7b377;
				text-transform: uppercase;
				letter-spacing: 3px;
				font-size: 2.8rem;
				text-align: center;
				margin-bottom: 2rem;
				text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5),
					0 0 20px rgba(199, 179, 119, 0.3);
				font-family: 'Cinzel', serif;
				font-weight: 700;
				position: relative; /* Add this for positioning the underline */
				}

				h1::after {
					content: '';
					display: block;
					width: 100%; 
					height: 3px;
					background-color: #c7b377; /* Match the text color */
					margin: 0 auto;
					margin-top: 10px; /* Adjust spacing below the text */
				}

				.card-title {
					color: #c7b377;
					font-weight: 700;
					letter-spacing: 1px;
				}

				canvas {
					background-color: rgba(20, 20, 20, 0.95);
					border-radius: 5px;
					border: 1px solid #8e7c44;
				}

				/* Add hover effects */
				.card:hover {
					transform: translateY(-2px);
					transition: all 0.3s ease;
					box-shadow: 0 0 15px rgba(199, 179, 119, 0.3);
				}

				/* Custom scrollbar styling */
				.card-body::-webkit-scrollbar {
					width: 8px;
				}

				.card-body::-webkit-scrollbar-track {
					background: rgba(20, 20, 20, 0.95);
				}

				.card-body::-webkit-scrollbar-thumb {
					background-color: #8e7c44;
					border-radius: 4px;
					border: 1px solid rgba(20, 20, 20, 0.95);
				}
			</style>
		`),
		app.H1().Class("text-center p-2").Text("Elder Scrolls Online"),
		app.Div().Class("row mt-4 d-flex align-items-stretch").Body(
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
		app.Div().Class("row mt-4").Body(
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
		// Embed JavaScript
		app.Raw(`
			<script>
				const statuses = ["All servers operational", "Minor issues detected", "Server maintenance ongoing", "Critical failure detected"];
				function updateStatus() {
					fetch('api/v1/eso_servers')
						.then(response => response.json())
						.then(data => {
							const serverStatusList = document.getElementById('serverStatusList');
							serverStatusList.innerHTML = ''; // Clear the list
							for (const [server, status] of Object.entries(data)) {
								let serversOffline = 0;
								if (status !== "Online") {
									serversOffline++
								}
								const listItem = document.createElement('li');
								listItem.className = 'list-group-item text-white';
								listItem.style.backgroundColor = status === 'Online' ? 'green' : 'red';
								listItem.textContent = server + " : " + status;
								serverStatusList.appendChild(listItem);

								switch (serversOffline) {
									case 0:
										document.getElementById('status').innerText = statuses[0];
										break;
									case 1:
										document.getElementById('status').innerText = statuses[1];
										break;
									default:
										document.getElementById('status').innerText = statuses[3];
								}
							}
						})
						.catch(error => {
							console.error('Error fetching server statuses:', error);
							const serverStatusList = document.getElementById('serverStatusList');
							document.getElementById('status').innerText = statuses[3];
							serverStatusList.innerHTML = '<li class="list-group-item text-white bg-danger">Error loading server statuses</li>';
						});
				}

				function updateStats() {
					fetch('api/v1/eso_current_players')
						.then(response => response.json())
						.then(data => {
							document.getElementById('activeUsers').innerText = data;
						})
						.catch(error => {
							console.error('Error fetching active users:', error);
							document.getElementById('activeUsers').innerText = 'Error';
						});

					fetch('api/v1/steam_charts_player_count')
						.then(response => response.json())
						.then(data => {
							if (data["24-peak"] !== undefined) {
								document.getElementById('24Peak').innerText = data["24-peak"];
							}
							if (data["all-time-peak"] !== undefined) {
								document.getElementById('allPeak').innerText = data["all-time-peak"];
							}
						})
						.catch(error => {
							console.error('Error fetching active users:', error);
							document.getElementById('24Peak').innerText = 'Error';
							document.getElementById('allPeak').innerText = 'Error';
						});
				}

				function fetchRSSFeed() {
					fetch('https://api.rss2json.com/v1/api.json?rss_url=https://eso-hub.com/en/news/feed.rss')
					.then(response => response.json())
					.then(data => {
						const rssFeed = document.getElementById('rssFeed');
						rssFeed.innerHTML = ''; // Clear the list
						data.items.slice(0, 3).forEach(item => {
						const feedItem
                        feedItem.href = item.link;
                        feedItem.target = '_blank';
                        feedItem.className = 'relative block bg-gray-900 border border-gray-900 shadow-lg rounded overflow-hidden';
                        feedItem.style.maxWidth = '100%';
                        feedItem.style.textDecoration = 'none'; // Remove underline
                        feedItem.innerHTML = '
                                <img src="\${item.thumbnail}" alt="\${item.title}" width="100%">
                                <div class="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black via-black/60 to-transparent">
                                <h3 class="text-white text-2xl font-semibold shadow-black line-clamp-2" style="text-shadow: black 2px 2px 1px;">
                                    \${item.title}
                                </h3>
                                <p class="text-white mt-2 line-clamp-2 text-sm">
                                    \${item.description}
                                </p>
                                <p class="text-white text-sm mt-2 line-clamp-1">
                                    Published on <span class="font-semibold text-white">\${new Date(item.pubDate).toLocaleDateString()}</span>
                                </p>
                                </div>
                            ';
                        rssFeed.appendChild(feedItem);
                        });
                    })
                    .catch(error => {
                        console.error('Error fetching RSS feed:', error);
                        const rssFeed = document.getElementById('rssFeed');
                        rssFeed.innerHTML = '<div class="text-white bg-danger p-3 rounded">Error loading RSS feed</div>';
                    });
                }

                // Fetch RSS feed on page load
                fetchRSSFeed();

                // Automatically refresh status, stats, and chart every 60 seconds
                setInterval(() => {
                    updateStatus();
                    updateStats();
                }, 60000);

                // Initial updates on page load
                updateStatus();
                updateStats();
            </script>
        `),
	)
}
