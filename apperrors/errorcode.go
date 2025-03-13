package apperrors

type ErrCode string

const (

	/*
		Unknown is used when none of the codes defined below apply.
	*/
	Unknown ErrCode = "U000"

	/*
		Prefix S codes are used in the service layer
	*/
	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	/*
		Prefix R codes are used in the service layer
	*/
	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
