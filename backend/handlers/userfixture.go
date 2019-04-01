package handlers

import "net/http"

/*GetUserFixtures queries the database for the list of fixtures watched by a user.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetUserFixtures(w http.ResponseWriter, r *http.Request) {

}

/*CreateUserFixture is used to create or update a tbl_user_fixture row in the database.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func CreateUserFixture(w http.ResponseWriter, r *http.Request) {

}

/*DeleteUserFixture is used to delete a tbl_user_fixture row in the database.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: DELETE
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func DeleteUserFixture(w http.ResponseWriter, r *http.Request) {

}
