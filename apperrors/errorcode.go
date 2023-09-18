package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParameter        ErrCode = "R002"

	RequiredAuthorizationHeader ErrCode = "A001"
	CannotCreateValidator       ErrCode = "A002"
	Unauthorized                ErrCode = "A003"
	NotMatchUser                ErrCode = "A004"
)
