package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error
}

func (m *MyAppError) Error() string {
	return m.Err.Error()
}

func (m *MyAppError) unWrap() error {
	return m.Err
}
