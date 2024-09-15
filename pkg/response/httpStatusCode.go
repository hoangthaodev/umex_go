package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrInvalidToken     = 30001 // token is invalid
	// register code
	ErrCodeUserHasExists = 50001 // user has already registered
)

var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "User has already registered",
}
