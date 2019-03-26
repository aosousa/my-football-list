package handlers

import (
	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
)

/*GetUserByUsername queries for a User row in the database.
 * Receives:
 * username (string) - User's username
 *
 * Returns:
 * User - User struct found
 * err - Description of error found during execution (or nil otherwise)
 */
func GetUserByUsername(username string) (m.User, error) {
	var u m.User

	err := db.QueryRow("SELECT userId, username, password, email FROM tbl_user WHERE username = ?", username).Scan(&u.UserID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		utils.HandleError("User", "Single", err)
		return u, err
	}

	return u, nil
}
