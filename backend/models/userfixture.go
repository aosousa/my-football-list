package models

// UserFixture represents the struct related to the associative table
type UserFixture struct {
	UserFixtureID int     `json:"userFixtureId"` // Unique ID of the user's fixture record
	UserID        int     `json:"userId"`        // User ID
	Fixture       Fixture `json:"fixture"`     // Fixture struct
	Status        int     `json:"status"`        // User's relation with the fixture (1 = Watching, 2 = Watched, 3 = Interested in Watching)
}

// UserFixtures represents a slice of UserFixture structs
type UserFixtures []UserFixture

// UserFixtureRequest is the struct used to create a user fixture association through the POST endpoint
type UserFixtureRequest struct {
	UserFixtureID int `json:"userFixtureId"` // Unique ID of the user's fixture record (or nil when creating a new row instead of updating)
	Fixture       int `json:"fixtureId"`     // Fixture ID
	Status        int `json:"status"`        // User's relation with the fixture (1 = Watching, 2 = Watched, 3 = Interested in Watching)
}

// UserFixtureResponse is the struct used for a response in the /users/{id}/fixtures endpoint
type UserFixtureResponse struct {
	Watched 	 UserFixtures `json:"watched"` 
	InterestedIn UserFixtures `json:"interestedIn"`
}