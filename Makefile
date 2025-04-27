build:
	set GOARCH=wasm GOOS=js 
	go build -o web/app.wasm
	go build

run: build
	./go-eso-dashboard