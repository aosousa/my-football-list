package main

import "./models"

const (
	baseURL = "https://api-football-v1.p.rapidapi.com/"
	version = "0.0.0"
)

var config models.Config

func initConfig() (models.Config, error) {
	config, err := models.CreateConfig()
	return config, err
}
