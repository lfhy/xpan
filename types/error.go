package types

type Error struct {
	AuthError    string `json:"error"`
	AuthErrorMsg string `json:"error_description"`
	Errno        int    `json:"errno"`
	ErrMsg       string `json:"errmsg"`
	RequestId    string `json:"request_id"`
}

func (e Error) IsError() bool {
	return !(e.AuthError == "" && e.Errno == 0)
}

func (e Error) Error() string {
	if e.AuthErrorMsg != "" {
		return e.AuthErrorMsg
	}
	return e.ErrMsg
}
