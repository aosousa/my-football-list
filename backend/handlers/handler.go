package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
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
	var err error
	sqlStmt := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database)

	// establish connection to the database
	db, err = sql.Open("mysql", sqlStmt)
	if err != nil {
		utils.HandleError("Handler", "InitDatabase", err)
	}

	// because sql.Open does not actually open a connection, we need to validate DSN data
	err = db.Ping()
	if err != nil {
		utils.HandleError("Handler", "InitDatabase", err)
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

/*func SaveLeagues() {
	var (
		objMap   map[string]interface{}
		queryURL string
	)

	queryURL = baseURL + "leagues"
	fmt.Println(queryURL)

	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		utils.HandleError("Handler", "SaveLeagues", err)
	}

	req.Header.Set("X-RapidAPI-Key", config.APIKey)
	req.Header.Set("Accept", "application/json")

	// call RapidAPI
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != 200 || err != nil {
		fmt.Println("res")
		utils.HandleError("Handler", "SaveLeagues", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("content")
		utils.HandleError("Handler", "SaveLeagues", err)
	}

	if err := json.Unmarshal(content, &objMap); err != nil {
		fmt.Println("unmarshal")
		log.Fatal(err)
	}

	stmtIns, err := db.Prepare("INSERT INTO tbl_league (name, country, season, logoUrl, flagUrl) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("statementIns")
		utils.HandleError("Handler", "SaveLeagues", err)
	}
	defer stmtIns.Close()

	for i := 1; i < 404; i++ {
		var (
			leagueObj                                       map[string]interface{}
			leagueID, leagueName, country, logoUrl, flagUrl string
		)
		leagueID = strconv.Itoa(i)
		leagueObj = objMap["api"].(map[string]interface{})["leagues"].(map[string]interface{})[leagueID].(map[string]interface{})

		intSeason, _ := strconv.Atoi(leagueObj["season"].(string))

		leagueName, country, logoUrl, flagUrl = leagueObj["name"].(string), leagueObj["country"].(string), leagueObj["logo"].(string), leagueObj["flag"].(string)

		_, err = stmtIns.Exec(leagueName, country, intSeason, logoUrl, flagUrl)
		if err != nil {
			utils.HandleError("Handler", "SaveLeagues", err)
		}
	}
}*/

func SaveTeams(leagueID string) {
	var (
		objMap   map[string]interface{}
		queryURL string
	)

	queryURL = baseURL + "teams/league/" + leagueID

	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		utils.HandleError("Handler", "SaveTeams", err)
	}

	req.Header.Set("X-RapidAPI-Key", config.APIKey)
	req.Header.Set("Accept", "application/json")

	// call RapidAPI
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != 200 || err != nil {
		fmt.Println("res")
		utils.HandleError("Handler", "SaveTeams", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("content")
		utils.HandleError("Handler", "SaveTeams", err)
	}

	if err := json.Unmarshal(content, &objMap); err != nil {
		fmt.Println("unmarshal")
		log.Fatal(err)
	}

	teams := objMap["api"].(map[string]interface{})["teams"].(map[string]interface{})

	stmtIns, err := db.Prepare("INSERT INTO tbl_team (teamId, name, logoUrl) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println("statementIns")
		utils.HandleError("Handler", "SaveTeams", err)
	}
	defer stmtIns.Close()

	for k, v := range teams {
		var name, logoUrl string

		intID, _ := strconv.Atoi(k)
		name, logoUrl = v.(map[string]interface{})["name"].(string), v.(map[string]interface{})["logo"].(string)

		_, _ = stmtIns.Exec(intID, name, logoUrl)
	}
}

/*SetResponse sets the response to be sent to the user in any API endpoints.
 * Receives:
 * w (http.ResponseWriter) - Go's HTTP ResponseWriter struct
 * statusCode (int) - HTTP status code of the response
 * body (m.HTTPResponse) - Body of the HTTP response using a custom struct
 */
func SetResponse(w http.ResponseWriter, statusCode int, body models.HTTPResponse) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
	return
}
