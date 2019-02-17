package handlers

import (
	"fmt"

	"../models"
	"github.com/jasonlvhit/gocron"
)

const (
	baseURL = "https://api-football-v1.p.rapidapi.com/"
	version = "0.0.0"
)

var config models.Config

/*InitConfig adds information from a configuration file to a Config struct
 * that will be used throughout the application.
 */
func InitConfig() (models.Config, error) {
	config, _ = models.CreateConfig()
	return config, nil
}

/*StartCronJob starts a cron job that will update the database every 30 minutes
 * with new matches and also update the previously running matches.
 */
func StartCronJob() {
	scheduler := gocron.NewScheduler()
	scheduler.Every(config.RefreshTimer).Minutes().Do(updateMatches)
	<-scheduler.Start()
}

// Add new matches and update old ones in the database
func updateMatches() {
	fmt.Println("testing cron")
}
