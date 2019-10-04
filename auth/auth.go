package auth

// LoginWithExternalServer allows for a custom implementation of the
// authentication process. We'll send an HTTP request to the configured
// server, and the server must reply either the user ID or a failed login attempt.
// This is the most flexible authentication process in Blocks.
func LoginWithExternalServer() LoginAttempt {
	return LoginAttempt{
		success: false,
		token:   "",
	}
}

// LoginWithEmailAndPassword attempts to authenticate the user with email + password.
func LoginWithEmailAndPassword(email string, password string) LoginAttempt {
	return LoginAttempt{
		success: false,
		token:   "",
	}
}
