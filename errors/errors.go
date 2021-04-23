package errors

import "errors"

type ErrCode int16

const (
	Unknown ErrCode = -1 + iota
)

var errCodes = make(map[error]ErrCode)

func ErrorCode(err error) ErrCode {
	if code, ok := errCodes[err]; !ok {
		for _err, code := range errCodes {
			if errors.Is(err, _err) {
				return code
			}
		}
		return Unknown
	} else {
		return code
	}
}

func RegisterErrCode(err error, code ErrCode) {
	if _, ok := errCodes[err]; ok {
		return
	}

	errCodes[err] = code
}
