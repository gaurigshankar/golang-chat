package main

import (
  "net/http"

  "log"

  "strings"
  "strconv"
  "github.com/gaurigshankar/golang-chat/chat"
  "github.com/gaurigshankar/golang-chat/config"
)

var configuration config.Configuration
var serverHostName string

func init() {
  configuration = config.LoadConfigAndSetUpLogging()
  strs := []string{configuration.Hostname, ":", strconv.Itoa(configuration.Port)}
  serverHostName = strings.Join(strs, "")
  log.Println("The serverHost url", serverHostName)

}

func main() {

    // websocket server
	server := chat.NewServer()
	go server.Listen()
  http.HandleFunc("/messages", handleHomePage)
  http.HandleFunc("/", handleHomePage)

  http.ListenAndServe(serverHostName, nil)

}

func handleHomePage(responseWriter http.ResponseWriter, request *http.Request) {
  http.ServeFile(responseWriter, request, "chat.html")
}
