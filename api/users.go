package api

// User represents the user for this application
// swagger:model
type User struct {
	// the id for this user
	//
	// required: true
	ID string `json:"id"`
	// the email for this user
	//
	// required: true
	Email string `json:"email"`
	// the username for this user
	//
	// required: true
	Username string `json:"username"`
	// the password for this user
	//
	// required: true
	Password string `json:"password"`
	// the admin status for this user
	//
	// required: true
	Admin bool `json:"admin"`
}

var UsersDB = make(map[string]User)
