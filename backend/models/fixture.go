package models

/*Fixture represents a football match. Fields:
 * FixtureID (int) - Unique ID of the fixture
 * APIFixtureID (int) - Unique ID of the fixture received from the API
 * Timestamp (string) - Fixture start timestamp
 * League (League) - League struct with information about the fixture's league
 * Round (string) - League round in which the fixture is played
 * HomeTeam (Team) - Team struct with information about the fixture's home team
 * HomeTeamGoals (int) - Number of goals scored by the home team
 * AwayTeam (Team) - Team struct with information about the fixture's away team
 * AwayTeamGoals (int) - Number of goals scored by the away team
 * Status (string) - Status of the fixture (not started, ongoing, finished, etc.)
 * Elapsed (int) - Number of minutes played in the fixture
 */
type Fixture struct {
	FixtureID     int    `json:"fixtureId"`
	APIFixtureID  int    `json:"apiFixtureId"`
	Timestamp     string `json:"timestamp"`
	League        League `json:"league"`
	Round         string `json:"round"`
	HomeTeam      Team   `json:"homeTeam"`
	HomeTeamGoals int    `json:"homeTeamGoals"`
	AwayTeam      Team   `json:"awayTeam"`
	AwayTeamGoals int    `json:"awayTeamGoals"`
	Status        string `json:"status"`
	Elapsed       int    `json:"elapsed"`
}

// Fixtures represents a slice of Fixture structs
type Fixtures []Fixture
