package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerIP   string   `json:"serverIP"`
	ServerPort string   `json:"serverPort"`
	ImgWidth   int      `json:"imgWidth"`
	ImgHeigh   int      `json:"imgHeigh"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error: ", err)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
