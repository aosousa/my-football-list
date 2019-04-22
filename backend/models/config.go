package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/aosousa/my-football-list/utils"
)

//Config struct contains all the necessary configurations for the back-end to run
type Config struct {
	APIKey       string `json:"apiKey"`       // Key for football fixtures API
	RefreshTimer uint64 `json:"refreshTimer"` // Frequency (in minutes) of updates to the database
	DB           DB     `json:"database"`     // Database configuration
}

// DB struct contains the database configuration
type DB struct {
	Host     string `json:"host"`     // Database hostname or IP address
	Port     string `json:"port"`     // Port in which database is running
	User     string `json:"user"`     // User used to authenticate in the database
	Password string `json:"password"` // Password used to authenticate in the database
	Database string `json:"database"` // Name of database schema used
}

// CreateConfig adds information from a configuration file to a Config struct.
func CreateConfig() Config {
	var config Config
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		utils.HandleError("Config", "CreateConfig", err)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		utils.HandleError("Config", "CreateConfig", err)
	}

	if config.APIKey == "" {
		err = errors.New("API key missing")
		utils.HandleError("Config", "CreateConfig", err)
	}

	if config.RefreshTimer == 0 {
		err = errors.New("Refresh timer missing")
		utils.HandleError("Config", "CreateConfig", err)
	}

	if (DB{} == config.DB) {
		err = errors.New("Database configuration missing")
		utils.HandleError("Config", "CreateConfig", err)
	}

	return config
}
