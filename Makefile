default:
	go build -o home-network-monitor-macos
	clear
	./home-network-monitor-macos
arm-build:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o home-network-monitor-arm