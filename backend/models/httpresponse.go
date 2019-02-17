package models

/*HTTPResponse is the struct used to send back JSON responses in the API. Contains:
 * Success (bool) - Whether the request was successful or not
 * Rows (int) - Number of rows in the Data interface response
 * Data (interface{}) - Interface that has the relevant information to return. Can be a slice, object, string, etc.
 * Error (string) - Error message (in case an error occurred)
 */
type HTTPResponse struct {
	Success bool        `json:"success"`
	Rows    int         `json:"int"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}
