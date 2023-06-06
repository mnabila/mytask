run:
	go run main.go

build:
	cd frontend && yarn build
	CGO_ENABLED=0 go build -ldflags="-s -w -X main.RELEASE=YES" -o mytask
