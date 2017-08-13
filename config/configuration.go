package config

import (
  "os"
  "fmt"
  "log"
  "encoding/json"
)

type Configuration struct {
  Port int `json:"port"`
  Hostname string `json:"hostName"`
  LogfilePath string `json:"logfilePath"`
}


func loadConfig() Configuration {
  configuration := Configuration{}
  configFile, err := os.Open("./config/config.json")
  defer configFile.Close()
  if err != nil {
    fmt.Println("Error loading config.json",err.Error())
    panic("Error loading config.json")
  }
  decoder := json.NewDecoder(configFile)
  err = decoder.Decode(&configuration)
  if err != nil {
    panic("Error loading config object from config.json")
  }
  return configuration
}

func setupLogging(configuration Configuration) {
  nf, err := os.Create(configuration.LogfilePath)
  if err != nil {
    fmt.Println(err)
  }
  log.SetOutput(nf)
}

func LoadConfigAndSetUpLogging() Configuration {
  configuration := loadConfig()
  setupLogging(configuration)

  return configuration
}
