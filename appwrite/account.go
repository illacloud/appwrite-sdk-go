package appwrite

import (
	"strings"
)

// Account service
type Account struct {
	client Client
}

func NewAccount(clt Client) *Account {
	return &Account{
		client: clt,
	}
}

// Get get currently logged in user data as JSON object.
func (srv *Account) Get() (*ClientResponse, error) {
	path := "/account"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateEmail update currently logged in user account email address. After
// changing user address, the user confirmation status will get reset. A new
// confirmation email is not sent automatically however you can use the send
// confirmation email endpoint again to send the confirmation email. For
// security measures, user password is required to complete this request.
// This endpoint can also be used to convert an anonymous account to a normal
// one, by passing an email address and a new password.
// 
func (srv *Account) UpdateEmail(Email string, Password string) (*ClientResponse, error) {
	path := "/account/email"
	
	params := map[string]interface{}{
		"email": Email,
		"password": Password,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// ListLogs get currently logged in user list of latest security activity
// logs. Each log returns user IP address, location and date and time of log.
func (srv *Account) ListLogs(Queries []interface{}) (*ClientResponse, error) {
	path := "/account/logs"
	
	params := map[string]interface{}{
		"queries": Queries,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateName update currently logged in user account name.
func (srv *Account) UpdateName(Name string) (*ClientResponse, error) {
	path := "/account/name"
	
	params := map[string]interface{}{
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdatePassword update currently logged in user password. For validation,
// user is required to pass in the new password, and the old password. For
// users created with OAuth, Team Invites and Magic URL, oldPassword is
// optional.
func (srv *Account) UpdatePassword(Password string, OldPassword string) (*ClientResponse, error) {
	path := "/account/password"
	
	params := map[string]interface{}{
		"password": Password,
		"oldPassword": OldPassword,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdatePhone update the currently logged in user's phone number. After
// updating the phone number, the phone verification status will be reset. A
// confirmation SMS is not sent automatically, however you can use the [POST
// /account/verification/phone](/docs/client/account#accountCreatePhoneVerification)
// endpoint to send a confirmation SMS.
func (srv *Account) UpdatePhone(Phone string, Password string) (*ClientResponse, error) {
	path := "/account/phone"
	
	params := map[string]interface{}{
		"phone": Phone,
		"password": Password,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// GetPrefs get currently logged in user preferences as a key-value object.
func (srv *Account) GetPrefs() (*ClientResponse, error) {
	path := "/account/prefs"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdatePrefs update currently logged in user account preferences. The object
// you pass is stored as is, and replaces any previous value. The maximum
// allowed prefs size is 64kB and throws error if exceeded.
func (srv *Account) UpdatePrefs(Prefs interface{}) (*ClientResponse, error) {
	path := "/account/prefs"
	
	params := map[string]interface{}{
		"prefs": Prefs,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// CreateRecovery sends the user an email with a temporary secret key for
// password reset. When the user clicks the confirmation link he is redirected
// back to your app password reset URL with the secret key and email address
// values attached to the URL query string. Use the query string params to
// submit a request to the [PUT
// /account/recovery](/docs/client/account#accountUpdateRecovery) endpoint to
// complete the process. The verification link sent to the user's email
// address is valid for 1 hour.
func (srv *Account) CreateRecovery(Email string, Url string) (*ClientResponse, error) {
	path := "/account/recovery"
	
	params := map[string]interface{}{
		"email": Email,
		"url": Url,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// UpdateRecovery use this endpoint to complete the user account password
// reset. Both the **userId** and **secret** arguments will be passed as query
// parameters to the redirect URL you have provided when sending your request
// to the [POST /account/recovery](/docs/client/account#accountCreateRecovery)
// endpoint.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
func (srv *Account) UpdateRecovery(UserId string, Secret string, Password string, PasswordAgain string) (*ClientResponse, error) {
	path := "/account/recovery"
	
	params := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
		"password": Password,
		"passwordAgain": PasswordAgain,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// ListSessions get currently logged in user list of active sessions across
// different devices.
func (srv *Account) ListSessions() (*ClientResponse, error) {
	path := "/account/sessions"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteSessions delete all sessions from the user account and remove any
// sessions cookies from the end client.
func (srv *Account) DeleteSessions() (*ClientResponse, error) {
	path := "/account/sessions"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// GetSession use this endpoint to get a logged in user's session using a
// Session ID. Inputting 'current' will return the current session being used.
func (srv *Account) GetSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateSession access tokens have limited lifespan and expire to mitigate
// security risks. If session was created using an OAuth provider, this route
// can be used to "refresh" the access token.
func (srv *Account) UpdateSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// DeleteSession use this endpoint to log out the currently logged in user
// from all their account sessions across all of their different devices. When
// using the Session ID argument, only the unique session ID provided is
// deleted.
// 
func (srv *Account) DeleteSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateStatus block the currently logged in user account. Behind the scene,
// the user record is not deleted but permanently blocked from any access. To
// completely delete a user, use the Users API instead.
func (srv *Account) UpdateStatus() (*ClientResponse, error) {
	path := "/account/status"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// CreateVerification use this endpoint to send a verification message to your
// user email address to confirm they are the valid owners of that address.
// Both the **userId** and **secret** arguments will be passed as query
// parameters to the URL you have provided to be attached to the verification
// email. The provided URL should redirect the user back to your app and allow
// you to complete the verification process by verifying both the **userId**
// and **secret** parameters. Learn more about how to [complete the
// verification process](/docs/client/account#accountUpdateEmailVerification).
// The verification link sent to the user's email address is valid for 7 days.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md),
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
// 
func (srv *Account) CreateVerification(Url string) (*ClientResponse, error) {
	path := "/account/verification"
	
	params := map[string]interface{}{
		"url": Url,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// UpdateVerification use this endpoint to complete the user email
// verification process. Use both the **userId** and **secret** parameters
// that were attached to your app URL to verify the user email ownership. If
// confirmed this route will return a 200 status code.
func (srv *Account) UpdateVerification(UserId string, Secret string) (*ClientResponse, error) {
	path := "/account/verification"
	
	params := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// CreatePhoneVerification use this endpoint to send a verification SMS to the
// currently logged in user. This endpoint is meant for use after updating a
// user's phone number using the
// [accountUpdatePhone](/docs/client/account#accountUpdatePhone) endpoint.
// Learn more about how to [complete the verification
// process](/docs/client/account#accountUpdatePhoneVerification). The
// verification code sent to the user's phone number is valid for 15 minutes.
func (srv *Account) CreatePhoneVerification() (*ClientResponse, error) {
	path := "/account/verification/phone"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// UpdatePhoneVerification use this endpoint to complete the user phone
// verification process. Use the **userId** and **secret** that were sent to
// your user's phone number to verify the user email ownership. If confirmed
// this route will return a 200 status code.
func (srv *Account) UpdatePhoneVerification(UserId string, Secret string) (*ClientResponse, error) {
	path := "/account/verification/phone"
	
	params := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}
