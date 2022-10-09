package apperrors

type ErrCode string

const (
	Unknown          ErrCode = "U000"
	InsertDataFailed ErrCode = "S001" // データベースへの insert処理に失敗
	GetDataFailed    ErrCode = "S002" // select文の実行に失敗
	NAData           ErrCode = "S003" // 指定された記事がない
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001" // リクエストボディの jsonデコードに失敗
	BadParam            ErrCode = "R002" // リクエストに含まれているパラメータが不正
)
