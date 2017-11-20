
mkdir $(pwd)/go
GOPATH="$(pwd)/go"
go get github.com/eclipse/paho.mqtt.golang
go get golang.org/x/net/websocket
go get golang.org/x/net/proxy

