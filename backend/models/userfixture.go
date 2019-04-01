package models

/*UserFixture represents the struct related to the associative table
 * used to store a user's fixtures records. Fields:
 * UserFixtureID (int) - Unique ID of the user's fixture record
 * User (User) - User struct
 * Fixture (Fixture) - Fixture struct
 * Status (int) - Status of the user's relation with the fixture (1 = Watching, 2 = Watched, 3 = Interested in Watching)
 */
type UserFixture struct {
	UserFixtureID int     `json:"userFixtureId"`
	User          User    `json:"userId"`
	Fixture       Fixture `json:"fixtureId"`
	Status        int     `json:"status"`
}

// UserFixtures represents a slice of UserFixture structs
type UserFixtures []UserFixture
