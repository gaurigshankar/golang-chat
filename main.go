package main

import (
  "net/http"
  "os"
  "log"
  "fmt"
  "strconv"
  "github.com/gaurigshankar/golang-chat/chat"
  "github.com/gaurigshankar/golang-chat/config"
)

var configuration config.Configuration
var serverHostName string

func init() {
  configuration = config.LoadConfigAndSetUpLogging()

  herokuConfigPort := os.Getenv("PORT")
  if herokuConfigPort == "" {
		serverHostName = fmt.Sprintf("%s:%s",configuration.Hostname,strconv.Itoa(configuration.Port))
	} else {
    serverHostName = fmt.Sprintf(":%s",herokuConfigPort)
  }
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
