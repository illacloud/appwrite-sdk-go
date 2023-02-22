package appwrite

import (
	"strings"
)

// Users service
type Users struct {
	client Client
}

func NewUsers(clt Client) *Users {
	return &Users{
		client: clt,
	}
}

// List get a list of all the project's users. You can use the query params to
// filter your results.
func (srv *Users) List(Queries []interface{}, Search string) (*ClientResponse, error) {
	path := "/users"
	
	params := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Create create a new user.
func (srv *Users) Create(UserId string, Email string, Phone string, Password string, Name string) (*ClientResponse, error) {
	path := "/users"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"phone": Phone,
		"password": Password,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateArgon2User create a new user. Password provided must be hashed with
// the [Argon2](https://en.wikipedia.org/wiki/Argon2) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateArgon2User(UserId string, Email string, Password string, Name string) (*ClientResponse, error) {
	path := "/users/argon2"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateBcryptUser create a new user. Password provided must be hashed with
// the [Bcrypt](https://en.wikipedia.org/wiki/Bcrypt) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateBcryptUser(UserId string, Email string, Password string, Name string) (*ClientResponse, error) {
	path := "/users/bcrypt"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateMD5User create a new user. Password provided must be hashed with the
// [MD5](https://en.wikipedia.org/wiki/MD5) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateMD5User(UserId string, Email string, Password string, Name string) (*ClientResponse, error) {
	path := "/users/md5"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreatePHPassUser create a new user. Password provided must be hashed with
// the [PHPass](https://www.openwall.com/phpass/) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreatePHPassUser(UserId string, Email string, Password string, Name string) (*ClientResponse, error) {
	path := "/users/phpass"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateScryptUser create a new user. Password provided must be hashed with
// the [Scrypt](https://github.com/Tarsnap/scrypt) algorithm. Use the [POST
// /users](/docs/server/users#usersCreate) endpoint to create users with a
// plain text password.
func (srv *Users) CreateScryptUser(UserId string, Email string, Password string, PasswordSalt string, PasswordCpu int, PasswordMemory int, PasswordParallel int, PasswordLength int, Name string) (*ClientResponse, error) {
	path := "/users/scrypt"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"passwordSalt": PasswordSalt,
		"passwordCpu": PasswordCpu,
		"passwordMemory": PasswordMemory,
		"passwordParallel": PasswordParallel,
		"passwordLength": PasswordLength,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateScryptModifiedUser create a new user. Password provided must be
// hashed with the [Scrypt
// Modified](https://gist.github.com/Meldiron/eecf84a0225eccb5a378d45bb27462cc)
// algorithm. Use the [POST /users](/docs/server/users#usersCreate) endpoint
// to create users with a plain text password.
func (srv *Users) CreateScryptModifiedUser(UserId string, Email string, Password string, PasswordSalt string, PasswordSaltSeparator string, PasswordSignerKey string, Name string) (*ClientResponse, error) {
	path := "/users/scrypt-modified"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"passwordSalt": PasswordSalt,
		"passwordSaltSeparator": PasswordSaltSeparator,
		"passwordSignerKey": PasswordSignerKey,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateSHAUser create a new user. Password provided must be hashed with the
// [SHA](https://en.wikipedia.org/wiki/Secure_Hash_Algorithm) algorithm. Use
// the [POST /users](/docs/server/users#usersCreate) endpoint to create users
// with a plain text password.
func (srv *Users) CreateSHAUser(UserId string, Email string, Password string, PasswordVersion string, Name string) (*ClientResponse, error) {
	path := "/users/sha"
	
	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"passwordVersion": PasswordVersion,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Get get a user by its unique ID.
func (srv *Users) Get(UserId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Delete delete a user by its unique ID, thereby releasing it's ID. Since ID
// is released and can be reused, all user-related resources like documents or
// storage files should be deleted before user deletion. If you want to keep
// ID reserved, use the [updateStatus](/docs/server/users#usersUpdateStatus)
// endpoint instead.
func (srv *Users) Delete(UserId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateEmail update the user email by its unique ID.
func (srv *Users) UpdateEmail(UserId string, Email string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/email")

	params := map[string]interface{}{
		"email": Email,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// ListLogs get the user activity logs list by its unique ID.
func (srv *Users) ListLogs(UserId string, Queries []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/logs")

	params := map[string]interface{}{
		"queries": Queries,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListMemberships get the user membership list by its unique ID.
func (srv *Users) ListMemberships(UserId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/memberships")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateName update the user name by its unique ID.
func (srv *Users) UpdateName(UserId string, Name string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/name")

	params := map[string]interface{}{
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdatePassword update the user password by its unique ID.
func (srv *Users) UpdatePassword(UserId string, Password string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/password")

	params := map[string]interface{}{
		"password": Password,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdatePhone update the user phone by its unique ID.
func (srv *Users) UpdatePhone(UserId string, Number string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/phone")

	params := map[string]interface{}{
		"number": Number,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// GetPrefs get the user preferences by its unique ID.
func (srv *Users) GetPrefs(UserId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdatePrefs update the user preferences by its unique ID. The object you
// pass is stored as is, and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws error if exceeded.
func (srv *Users) UpdatePrefs(UserId string, Prefs interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{
		"prefs": Prefs,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// ListSessions get the user sessions list by its unique ID.
func (srv *Users) ListSessions(UserId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteSessions delete all user's sessions by using the user's unique ID.
func (srv *Users) DeleteSessions(UserId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// DeleteSession delete a user sessions by its unique ID.
func (srv *Users) DeleteSession(UserId string, SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId, "{sessionId}", SessionId)
	path := r.Replace("/users/{userId}/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateStatus update the user status by its unique ID. Use this endpoint as
// an alternative to deleting a user if you want to keep user's ID reserved.
func (srv *Users) UpdateStatus(UserId string, Status bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/status")

	params := map[string]interface{}{
		"status": Status,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdateEmailVerification update the user email verification status by its
// unique ID.
func (srv *Users) UpdateEmailVerification(UserId string, EmailVerification bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/verification")

	params := map[string]interface{}{
		"emailVerification": EmailVerification,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdatePhoneVerification update the user phone verification status by its
// unique ID.
func (srv *Users) UpdatePhoneVerification(UserId string, PhoneVerification bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/verification/phone")

	params := map[string]interface{}{
		"phoneVerification": PhoneVerification,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}
