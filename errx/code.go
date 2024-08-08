package errx

import "google.golang.org/grpc/codes"

const (
	OK                 = uint32(codes.OK)
	InvalidArgument    = uint32(codes.InvalidArgument)
	Canceled           = uint32(codes.Canceled)
	Unknown            = uint32(codes.Unknown)
	DeadlineExceeded   = uint32(codes.DeadlineExceeded)
	NotFound           = uint32(codes.NotFound)
	AlreadyExists      = uint32(codes.AlreadyExists)
	PermissionDenied   = uint32(codes.PermissionDenied)
	ResourceExhausted  = uint32(codes.ResourceExhausted)
	FailedPrecondition = uint32(codes.FailedPrecondition)
	Aborted            = uint32(codes.Aborted)
	OutOfRange         = uint32(codes.OutOfRange)
	Unimplemented      = uint32(codes.Unimplemented)
	Internal           = uint32(codes.Internal)
	Unavailable        = uint32(codes.Unavailable)
	DataLoss           = uint32(codes.DataLoss)
	Unauthenticated    = uint32(codes.Unauthenticated)
)

const (
	CodeTokenExpireError   uint32 = 100101
	CodeTokenGenerateError uint32 = 100102
	CodeFailure            uint32 = 100103
	CodeNoLogin            uint32 = 100104
	CodeMethodNotAllow     uint32 = 100105
	CodeInternalServerErr  uint32 = 100106
)

const (
	CodeDbError                   uint32 = 100201
	CodeDbInsertErr               uint32 = 100202
	CodeDbUpdateAffectedZeroError uint32 = 100203
	CodeDbNotFoundErr             uint32 = 100204
)

var CodeToMsg = map[uint32]string{
	OK:                 MsgOk,
	Canceled:           MsgCanceled,
	Unknown:            MsgUnknown,
	InvalidArgument:    MsgInvalidArgument,
	DeadlineExceeded:   MsgDeadlineExceeded,
	NotFound:           MsgNotFound,
	AlreadyExists:      MsgAlreadyExists,
	PermissionDenied:   MsgPermissionDenied,
	ResourceExhausted:  MsgResourceExhausted,
	FailedPrecondition: MsgFailedPrecondition,
	Aborted:            MsgAborted,
	OutOfRange:         MsgOutOfRange,
	Unimplemented:      MsgUnimplemented,
	Internal:           MsgInternal,
	Unavailable:        MsgUnavailable,
	DataLoss:           MsgDataLoss,
	Unauthenticated:    MsgUnauthenticated,

	CodeTokenExpireError:          MsgTokenExpireError,
	CodeTokenGenerateError:        MsgTokenGenerateError,
	CodeDbError:                   MsgDbError,
	CodeDbUpdateAffectedZeroError: MsgDbUpdateAffectedZeroError,
	CodeFailure:                   MsgFailure,
	CodeNoLogin:                   MsgNoLogin,
	CodeMethodNotAllow:            MsgMethodNotAllow,
	CodeInternalServerErr:         MsgInternalServerErr,
}
