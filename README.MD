# Go-App ESO Dashboard

Simple Go-App PWA ESO dashboard with caching support for local deployment

![Project Overview](readme.png)

## Features

- ESO player count, server status and RSS feed from ESO hub.
- Easy to extend and customize.

## Requirements

- Go 1.24.1

## Installation

1. Clone the repository:
    ```bash
    git clone <repository-url>
    cd go-eso-dashboard
    ```

2. Install dependencies:
    ```bash
    go mod vendor
    ```

## Usage

1. Run the application:
    ```bash
        go build ./cmd/go-eso-dashboard # For binary
        go build -o ../../web/app.wasm # You need to set GOOS=js and GOARCH=wasm for this to work
    ```

2. Open your browser and navigate to `http://127.0.0.1:8000`.

## Project Structure

Here is an overview of the project structure:

```
go-eso-dashboard/
├── WIP ----
├── LICENSE                # MIT license
├── README.md              # Project documentation
└── readme.png             # Project overview
```

This structure ensures clarity and ease of navigation for developers.

## License

This project is licensed under the MIT License.

## Legal

Elder Scrolls Online, SteamCharts, Steam, and all associated logos and designs are trademarks or registered trademarks of their respective owners.  
This project is not affiliated with ZeniMax Online Studios, Bethesda Softworks, Steam, or any of their family of sites.
