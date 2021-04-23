package errors

import "errors"

type ErrCode int16

const (
	ErrUnknown ErrCode = -1 + iota
	ErrAuto
)

var (
	errCodes            = make(map[error]ErrCode)
	lastErrCode ErrCode = 10000
)

func ErrorCode(err error) ErrCode {
	if code, ok := errCodes[err]; !ok {
		for _err, code := range errCodes {
			if errors.Is(err, _err) {
				return code
			}
		}
		return ErrUnknown
	} else {
		return code
	}
}

func RegisterErrCode(err error, code ErrCode) {
	if _, ok := errCodes[err]; ok {
		return
	}

	if code == ErrAuto {
		lastErrCode++
		code = lastErrCode
	} else {
		if code > lastErrCode {
			lastErrCode = code
		}
	}
	errCodes[err] = code
}

func New(s string) error {
	return errors.New(s)
}
