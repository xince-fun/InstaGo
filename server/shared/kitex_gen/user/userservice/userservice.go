// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"RegisterPhone": kitex.NewMethodInfo(
		registerPhoneHandler,
		newUserServiceRegisterPhoneArgs,
		newUserServiceRegisterPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RegisterEmail": kitex.NewMethodInfo(
		registerEmailHandler,
		newUserServiceRegisterEmailArgs,
		newUserServiceRegisterEmailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"LoginPhone": kitex.NewMethodInfo(
		loginPhoneHandler,
		newUserServiceLoginPhoneArgs,
		newUserServiceLoginPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"LoginEmail": kitex.NewMethodInfo(
		loginEmailHandler,
		newUserServiceLoginEmailArgs,
		newUserServiceLoginEmailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateEmail": kitex.NewMethodInfo(
		updateEmailHandler,
		newUserServiceUpdateEmailArgs,
		newUserServiceUpdateEmailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdatePhone": kitex.NewMethodInfo(
		updatePhoneHandler,
		newUserServiceUpdatePhoneArgs,
		newUserServiceUpdatePhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdatePasswd": kitex.NewMethodInfo(
		updatePasswdHandler,
		newUserServiceUpdatePasswdArgs,
		newUserServiceUpdatePasswdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateBirthDay": kitex.NewMethodInfo(
		updateBirthDayHandler,
		newUserServiceUpdateBirthDayArgs,
		newUserServiceUpdateBirthDayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadAvatar": kitex.NewMethodInfo(
		uploadAvatarHandler,
		newUserServiceUploadAvatarArgs,
		newUserServiceUploadAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateAvatarInfo": kitex.NewMethodInfo(
		updateAvatarInfoHandler,
		newUserServiceUpdateAvatarInfoArgs,
		newUserServiceUpdateAvatarInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetAvatar": kitex.NewMethodInfo(
		getAvatarHandler,
		newUserServiceGetAvatarArgs,
		newUserServiceGetAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterPhoneArgs)
	realResult := result.(*user.UserServiceRegisterPhoneResult)
	success, err := handler.(user.UserService).RegisterPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterPhoneArgs() interface{} {
	return user.NewUserServiceRegisterPhoneArgs()
}

func newUserServiceRegisterPhoneResult() interface{} {
	return user.NewUserServiceRegisterPhoneResult()
}

func registerEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterEmailArgs)
	realResult := result.(*user.UserServiceRegisterEmailResult)
	success, err := handler.(user.UserService).RegisterEmail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterEmailArgs() interface{} {
	return user.NewUserServiceRegisterEmailArgs()
}

func newUserServiceRegisterEmailResult() interface{} {
	return user.NewUserServiceRegisterEmailResult()
}

func loginPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginPhoneArgs)
	realResult := result.(*user.UserServiceLoginPhoneResult)
	success, err := handler.(user.UserService).LoginPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginPhoneArgs() interface{} {
	return user.NewUserServiceLoginPhoneArgs()
}

func newUserServiceLoginPhoneResult() interface{} {
	return user.NewUserServiceLoginPhoneResult()
}

func loginEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginEmailArgs)
	realResult := result.(*user.UserServiceLoginEmailResult)
	success, err := handler.(user.UserService).LoginEmail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginEmailArgs() interface{} {
	return user.NewUserServiceLoginEmailArgs()
}

func newUserServiceLoginEmailResult() interface{} {
	return user.NewUserServiceLoginEmailResult()
}

func updateEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateEmailArgs)
	realResult := result.(*user.UserServiceUpdateEmailResult)
	success, err := handler.(user.UserService).UpdateEmail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateEmailArgs() interface{} {
	return user.NewUserServiceUpdateEmailArgs()
}

func newUserServiceUpdateEmailResult() interface{} {
	return user.NewUserServiceUpdateEmailResult()
}

func updatePhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdatePhoneArgs)
	realResult := result.(*user.UserServiceUpdatePhoneResult)
	success, err := handler.(user.UserService).UpdatePhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdatePhoneArgs() interface{} {
	return user.NewUserServiceUpdatePhoneArgs()
}

func newUserServiceUpdatePhoneResult() interface{} {
	return user.NewUserServiceUpdatePhoneResult()
}

func updatePasswdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdatePasswdArgs)
	realResult := result.(*user.UserServiceUpdatePasswdResult)
	success, err := handler.(user.UserService).UpdatePasswd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdatePasswdArgs() interface{} {
	return user.NewUserServiceUpdatePasswdArgs()
}

func newUserServiceUpdatePasswdResult() interface{} {
	return user.NewUserServiceUpdatePasswdResult()
}

func updateBirthDayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateBirthDayArgs)
	realResult := result.(*user.UserServiceUpdateBirthDayResult)
	success, err := handler.(user.UserService).UpdateBirthDay(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateBirthDayArgs() interface{} {
	return user.NewUserServiceUpdateBirthDayArgs()
}

func newUserServiceUpdateBirthDayResult() interface{} {
	return user.NewUserServiceUpdateBirthDayResult()
}

func uploadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUploadAvatarArgs)
	realResult := result.(*user.UserServiceUploadAvatarResult)
	success, err := handler.(user.UserService).UploadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUploadAvatarArgs() interface{} {
	return user.NewUserServiceUploadAvatarArgs()
}

func newUserServiceUploadAvatarResult() interface{} {
	return user.NewUserServiceUploadAvatarResult()
}

func updateAvatarInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateAvatarInfoArgs)
	realResult := result.(*user.UserServiceUpdateAvatarInfoResult)
	success, err := handler.(user.UserService).UpdateAvatarInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateAvatarInfoArgs() interface{} {
	return user.NewUserServiceUpdateAvatarInfoArgs()
}

func newUserServiceUpdateAvatarInfoResult() interface{} {
	return user.NewUserServiceUpdateAvatarInfoResult()
}

func getAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetAvatarArgs)
	realResult := result.(*user.UserServiceGetAvatarResult)
	success, err := handler.(user.UserService).GetAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetAvatarArgs() interface{} {
	return user.NewUserServiceGetAvatarArgs()
}

func newUserServiceGetAvatarResult() interface{} {
	return user.NewUserServiceGetAvatarResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) RegisterPhone(ctx context.Context, req *user.RegisterPhoneRequest) (r *user.RegisterResponse, err error) {
	var _args user.UserServiceRegisterPhoneArgs
	_args.Req = req
	var _result user.UserServiceRegisterPhoneResult
	if err = p.c.Call(ctx, "RegisterPhone", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RegisterEmail(ctx context.Context, req *user.RegisterEmailRequest) (r *user.RegisterResponse, err error) {
	var _args user.UserServiceRegisterEmailArgs
	_args.Req = req
	var _result user.UserServiceRegisterEmailResult
	if err = p.c.Call(ctx, "RegisterEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LoginPhone(ctx context.Context, req *user.LoginPhoneRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceLoginPhoneArgs
	_args.Req = req
	var _result user.UserServiceLoginPhoneResult
	if err = p.c.Call(ctx, "LoginPhone", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LoginEmail(ctx context.Context, req *user.LoginEmailRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceLoginEmailArgs
	_args.Req = req
	var _result user.UserServiceLoginEmailResult
	if err = p.c.Call(ctx, "LoginEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateEmail(ctx context.Context, req *user.UpdateEmailRequest) (r *user.UpdateEmailResponse, err error) {
	var _args user.UserServiceUpdateEmailArgs
	_args.Req = req
	var _result user.UserServiceUpdateEmailResult
	if err = p.c.Call(ctx, "UpdateEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePhone(ctx context.Context, req *user.UpdatePhoneRequest) (r *user.UpdatePhoneResponse, err error) {
	var _args user.UserServiceUpdatePhoneArgs
	_args.Req = req
	var _result user.UserServiceUpdatePhoneResult
	if err = p.c.Call(ctx, "UpdatePhone", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePasswd(ctx context.Context, req *user.UpdatePasswdRequest) (r *user.UpdatePasswdResponse, err error) {
	var _args user.UserServiceUpdatePasswdArgs
	_args.Req = req
	var _result user.UserServiceUpdatePasswdResult
	if err = p.c.Call(ctx, "UpdatePasswd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBirthDay(ctx context.Context, req *user.UpdateBirthDayRequest) (r *user.UpdateBirthDayResponse, err error) {
	var _args user.UserServiceUpdateBirthDayArgs
	_args.Req = req
	var _result user.UserServiceUpdateBirthDayResult
	if err = p.c.Call(ctx, "UpdateBirthDay", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (r *user.UploadAvatarResponse, err error) {
	var _args user.UserServiceUploadAvatarArgs
	_args.Req = req
	var _result user.UserServiceUploadAvatarResult
	if err = p.c.Call(ctx, "UploadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAvatarInfo(ctx context.Context, req *user.UpdateAvatarInfoRequest) (r *user.UpdateAvatarInfoResponse, err error) {
	var _args user.UserServiceUpdateAvatarInfoArgs
	_args.Req = req
	var _result user.UserServiceUpdateAvatarInfoResult
	if err = p.c.Call(ctx, "UpdateAvatarInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAvatar(ctx context.Context, req *user.GetAvatarRequest) (r *user.GetAvatarResponse, err error) {
	var _args user.UserServiceGetAvatarArgs
	_args.Req = req
	var _result user.UserServiceGetAvatarResult
	if err = p.c.Call(ctx, "GetAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
