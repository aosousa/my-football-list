package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"../utils"
)

/*Config struct contains all the necessary configurations for the back-end
to run. Contains:
 * APIKey (string) - Key for the football API
 * RefreshTimer (int) - Frequency (in minutes) of updates to the database
 * DB (DB) - Database configuration
*/
type Config struct {
	APIKey       string `json:"apiKey"`
	RefreshTimer uint64 `json:"refreshTimer"`
	DB           DB     `json:"database"`
}

// DB struct contains the database configuration
type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
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
