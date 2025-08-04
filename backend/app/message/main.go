package message

var (
	Success             = "success"
	InternalServerError = "internal-server-error"
	Forbidden           = "forbidden"
	Unauthorized        = "unauthorized"
	InvalidRequest      = "invalid-request-form"

	UserAlreadyExists = "user-already-exists"
	UserNotFound      = "user-not-found"
	UserIsInUse       = "user-is-in-use"

	EmailAlreadyExists = "email-already-exists"
	EmailNotFound      = "email-not-found"
	EmailIsInUse       = "email-is-in-use"

	PasswordIncorrect = "password-incorrect"
	PasswordNotMatch  = "password-not-match"
	OTPExpired        = "otp-expired"
	OTPInvalid        = "otp-invalid"
	OTPAlreadyUsed    = "otp-already-used"
	OTPNotFound       = "otp-not-found"

	InvalidCredentials = "username-or-password-incorrect"
)
