package errors

const (
	BadRequestCode          GlobalErrorCode = "BAD_REQUEST"
	InternalServerErrorCode GlobalErrorCode = "INTERNAL_SERVER_ERROR"

	ConflictErrorCode     GlobalErrorCode = "CONFLICT"
	UserAlreadyExistsCode GlobalErrorCode = "USER_ALREADY_EXISTS"
)
