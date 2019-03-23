package models

/*User represents a user in the platform. Fields:
 * UserID (int) - Unique ID of the user in the platform
 * Username (string) - User's username in the platform
 * Password (string) - User's password in the platform
 * Email (string) - User's email (used for account confirmation)
 * CreateTime (string) - Timestamp stating when the account was created
 * UpdateTime (string) - Timestamp stating when the account was last updated
 * Status (int) - Status of the account (0 = Unconfirmed, 1 = Confirmed)
 */
type User struct {
	UserID     int    `json:"userId"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Status     int    `json:"status"`
}

// Users represents a slice of User structs
type Users []User
