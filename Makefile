setup:
	go get ./..

start:
	go run main.go
	ps aux | grep 'tor -socksport'

stop:
	killall tor

ps:
	ps aux | grep 'tor -socksport'