package models

/*User represents a user in the platform. Fields:
 * UserID (int) - Unique ID of the user in the platform
 * Username (string) - User's username in the platform
 * Password (string) - User's password in the platform
 * Email (string) - User's email (used for account confirmation)
 * CreateTime (string) - Timestamp stating when the account was created
 * UpdateTime (string) - Timestamp stating when the account was last updated
 * Status (int) - Status of the account (0 = Unconfirmed, 1 = Confirmed)
 * SpoilerMode (int) - Whether or not to show results for this user (0 = Show, 1 = Hide)
 */
type User struct {
	UserID      int    `json:"userId"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	Status      int    `json:"status"`
	SpoilerMode int    `json:"spoilerMode"`
}

// Users represents a slice of User structs
type Users []User
