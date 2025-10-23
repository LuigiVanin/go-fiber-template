package errors

import "fmt"

type GlobalErrorCode string

type GlobalError struct {
	Title  string
	Code   GlobalErrorCode
	Detail string
	Type   string
}

func NewGlobalError(title string, code GlobalErrorCode, detail string) *GlobalError {

	return &GlobalError{
		Title:  title,
		Code:   code,
		Detail: detail,
	}
}

func ThrowBadRequest(detail string) *GlobalError {
	return &GlobalError{
		Title:  "Bad Request",
		Code:   BadRequestCode,
		Detail: detail,
	}
}

func ThrowInternalServerError(detail string) *GlobalError {
	return &GlobalError{
		Title:  "Internal Server Error",
		Code:   InternalServerErrorCode,
		Detail: detail,
	}
}

func ThrowConflict(detail string) *GlobalError {
	return &GlobalError{
		Title:  "Conflict",
		Code:   ConflictErrorCode,
		Detail: detail,
	}
}

func ThrowUserAlreadyExists(detail string) *GlobalError {
	return &GlobalError{
		Title:  "Conflict",
		Code:   UserAlreadyExistsCode,
		Detail: detail,
		Type:   "https://example.com/errors/user-already-exists",
	}
}

func (e *GlobalError) Error() string {

	return fmt.Sprintf("GlobalError: %s, Code: %s, Detail: %s", e.Title, e.Code, e.Detail)
}
