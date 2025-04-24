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
                listItem.textContent = `${server}: ${status}`;
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
        const feedItem = document.createElement('a');
        feedItem.href = item.link;
        feedItem.target = '_blank';
        feedItem.className = 'relative block bg-gray-900 border border-gray-900 shadow-lg rounded overflow-hidden';
        feedItem.style.maxWidth = '100%';
        feedItem.style.textDecoration = 'none'; // Remove underline
        feedItem.innerHTML = `
                <img src="${item.thumbnail}" alt="${item.title}" width="100%">
                <div class="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black via-black/60 to-transparent">
                <h3 class="text-white text-2xl font-semibold shadow-black line-clamp-2" style="text-shadow: black 2px 2px 1px;">
                    ${item.title}
                </h3>
                <p class="text-white mt-2 line-clamp-2 text-sm">
                    ${item.description}
                </p>
                <p class="text-white text-sm mt-2 line-clamp-1">
                    Published on <span class="font-semibold text-white">${new Date(item.pubDate).toLocaleDateString()}</span>
                </p>
                </div>
            `;
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
    console.log("Refreshing status, stats, and chart...");
    updateStatus();
    updateStats();
}, 60000);

// Initial updates on page load
updateStatus();
updateStats();