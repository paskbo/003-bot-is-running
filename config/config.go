package config

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

var (
  ChvfefeToken string
  AwesomePrefix string

  config *configStruct
)

type configStruct struct {
  ChvfefeToken string `json:"ChvfefeToken"`
  AwesomePrefix string `json: "AwesomePrefix"`
}

func ReadConfig() error {
  fmt.Println("The chvfefe is busy cooking the config files, it needs some salt and pepper...")

  file, err := ioutil.ReadFile("./config.json")

  if err != nil {
    fmt.Println(err.Error())
    return err
  }
  fmt.Println(string(file))

  err = json.Unmarshal(file, &config)

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  ChvfefeToken = config.ChvfefeToken
  AwesomePrefix = config.AwesomePrefix

  return nil
}
