package handlers

import (
	"database/sql"
	"fmt"

	"github.com/aosousa/my-football-list/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonlvhit/gocron"
)

const (
	baseURL = "https://api-football-v1.p.rapidapi.com/"
	version = "0.0.0"
)

var (
	config models.Config
	db     *sql.DB
)

/*InitConfig adds information from a configuration file to a Config struct
 * that will be used throughout the application.
 */
func InitConfig() {
	config = models.CreateConfig()
}

/*InitDatabase establishes a connection to the database with parameters in the
 * configuration struct that was previously loaded.
 */
func InitDatabase() {
	sqlStmt := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database)

	// establish connection to the database
	db, err := sql.Open("mysql", sqlStmt)
	defer db.Close()

	// because sql.Open does not actually open a connection, we need to validate DSN data
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
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
