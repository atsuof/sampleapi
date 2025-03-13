package apperrors

type MyAppError struct {
	ErrCode `json:"err-code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (m *MyAppError) Error() string {
	return m.Err.Error()
}

func (m *MyAppError) unWrap() error {
	return m.Err
}
