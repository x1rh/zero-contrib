package errx

const (
	MsgOk                 = "OK"
	MsgCanceled           = "CANCELLED"
	MsgUnknown            = "UNKNOWN"
	MsgInvalidArgument    = "INVALID_ARGUMENT"
	MsgDeadlineExceeded   = "DEADLINE_EXCEEDED"
	MsgNotFound           = "NOT_FOUND"
	MsgAlreadyExists      = "ALREADY_EXISTS"
	MsgPermissionDenied   = "PERMISSION_DENIED"
	MsgResourceExhausted  = "RESOURCE_EXHAUSTED"
	MsgFailedPrecondition = "FAILED_PRECONDITION"
	MsgAborted            = "ABORTED"
	MsgOutOfRange         = "OUT_OF_RANGE"
	MsgUnimplemented      = "UNIMPLEMENTED"
	MsgInternal           = "INTERNAL"
	MsgUnavailable        = "UNAVAILABLE"
	MsgDataLoss           = "DATA_LOSS"
	MsgUnauthenticated    = "UNAUTHENTICATED"
)

const (
	MsgTokenExpireError          = "token expire"
	MsgTokenGenerateError        = "token generation error"
	MsgDbError                   = "db error"
	MsgDbUpdateAffectedZeroError = "affected db rows is zero"
	MsgFailure                   = "fail"
	MsgNoLogin                   = "no login"
	MsgMethodNotAllow            = "method not allow"
	MsgInternalServerErr         = "internal server error"
)

var MsgToCode = map[string]uint32{
	MsgOk:                 OK,
	MsgCanceled:           Canceled,
	MsgUnknown:            Unknown,
	MsgInvalidArgument:    InvalidArgument,
	MsgDeadlineExceeded:   DeadlineExceeded,
	MsgNotFound:           NotFound,
	MsgAlreadyExists:      AlreadyExists,
	MsgPermissionDenied:   PermissionDenied,
	MsgResourceExhausted:  ResourceExhausted,
	MsgFailedPrecondition: FailedPrecondition,
	MsgAborted:            Aborted,
	MsgOutOfRange:         OutOfRange,
	MsgUnimplemented:      Unimplemented,
	MsgInternal:           Internal,
	MsgUnavailable:        Unavailable,
	MsgDataLoss:           DataLoss,
	MsgUnauthenticated:    Unauthenticated,
}

func MapErrMsg(code uint32) string {
	if msg, ok := CodeToMsg[code]; ok {
		return msg
	} else {
		return "service temporarily unavailable"
	}
}
