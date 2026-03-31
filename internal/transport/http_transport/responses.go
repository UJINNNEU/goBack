package http_transport

import "errors"

var (
	errInvalidUserId = errors.New("invalid user id")
)

type errResp struct {
	Err string `json:"error"`
}

func newErrResp(err error) errResp {
	return errResp{Err: err.Error()}
}
