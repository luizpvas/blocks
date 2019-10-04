package auth

// LoginAttempt is the return value from Login* functions.
type LoginAttempt struct {
	success bool
	token   string
}
