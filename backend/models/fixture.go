package models

//Fixture represents a football match.
type Fixture struct {
	FixtureID         int    `json:"fixtureId"`         // Unique ID of the fixture
	APIFixtureID      int    `json:"apiFixtureId"`      // Unique ID of the fixture received from the API
	Date              string `json:"date"`              // Fixture date
	League            League `json:"league"`            // League struct with the information about the fixture's league
	Round             string `json:"round"`             // League round in which the fixture is played
	HomeTeam          Team   `json:"homeTeam"`          // Team struct with information about the fixture's home team
	HomeTeamGoals     int    `json:"homeTeamGoals"`     // Number of goals scored by the home team
	AwayTeam          Team   `json:"awayTeam"`          // Team struct with information about the fixture's away team
	AwayTeamGoals     int    `json:"awayTeamGoals"`     // Number of goals scored by the away team
	Status            string `json:"status"`            // Status of the fixture (not started, ongoing, finished, etc.)
	Elapsed           int    `json:"elapsed"`           // Number of minutes played in the fixture
	UserFixtureStatus int    `json:"userFixtureStatus"` // User's relationship with the fixture (1 = Watching, 2 = Watched, 3 = Interested in Watching)
}

// Fixtures represents a slice of Fixture structs
type Fixtures []Fixture
