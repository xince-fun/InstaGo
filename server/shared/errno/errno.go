package errno

import (
	"fmt"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/errno"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// NewErrNo return ErrNo
func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success               = NewErrNo(int64(errno.Err_Success), "success")
	BadRequest            = NewErrNo(int64(errno.Err_BadRequest), "bad request")
	ParamsErr             = NewErrNo(int64(errno.Err_ParamsErr), "params error")
	ServiceErr            = NewErrNo(int64(errno.Err_ServiceErr), "service error")
	UserSrvError          = NewErrNo(int64(errno.Err_UserSrvErr), "user service error")
	UserPwdError          = NewErrNo(int64(errno.Err_UserPwdErr), "user password error")
	UserPwdSameError      = NewErrNo(int64(errno.Err_UserPwdSameErr), "new password is the same as the old password")
	UserNotExistError     = NewErrNo(int64(errno.Err_UserNotExistErr), "user not exist")
	BlobSrvError          = NewErrNo(int64(errno.Err_BlobSrvErr), "blob service error")
	RelationDBError       = NewErrNo(int64(errno.Err_RelationDBErr), "follow db error")
	RelationSrvError      = NewErrNo(int64(errno.Err_RelationSrvErr), "follow service error")
	RelationSelfError     = NewErrNo(int64(errno.Err_RelationSelfErr), "can't follow self")
	RelationExistError    = NewErrNo(int64(errno.Err_RelationExistErr), "already followed")
	RelationNotExistError = NewErrNo(int64(errno.Err_RelationNotExistErr), "follow not exist")
	RecordNotFound        = NewErrNo(int64(errno.Err_RecordNotFound), "record not found")
	RecordExist           = NewErrNo(int64(errno.Err_RecordExist), "record exist")
)

var (
	InvalidDate = NewErrNo(int64(errno.Err_InvalidDate), "invalid date")
)
