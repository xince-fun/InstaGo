// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RegisterPhone(ctx context.Context, req *user.RegisterPhoneRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	RegisterEmail(ctx context.Context, req *user.RegisterEmailRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	LoginPhone(ctx context.Context, req *user.LoginPhoneRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	LoginEmail(ctx context.Context, req *user.LoginEmailRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	UpdateEmail(ctx context.Context, req *user.UpdateEmailRequest, callOptions ...callopt.Option) (r *user.UpdateEmailResponse, err error)
	UpdatePhone(ctx context.Context, req *user.UpdatePhoneRequest, callOptions ...callopt.Option) (r *user.UpdatePhoneResponse, err error)
	UpdatePasswd(ctx context.Context, req *user.UpdatePasswdRequest, callOptions ...callopt.Option) (r *user.UpdatePasswdResponse, err error)
	UpdateBirthDay(ctx context.Context, req *user.UpdateBirthDayRequest, callOptions ...callopt.Option) (r *user.UpdateBirthDayResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) RegisterPhone(ctx context.Context, req *user.RegisterPhoneRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RegisterPhone(ctx, req)
}

func (p *kUserServiceClient) RegisterEmail(ctx context.Context, req *user.RegisterEmailRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RegisterEmail(ctx, req)
}

func (p *kUserServiceClient) LoginPhone(ctx context.Context, req *user.LoginPhoneRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginPhone(ctx, req)
}

func (p *kUserServiceClient) LoginEmail(ctx context.Context, req *user.LoginEmailRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginEmail(ctx, req)
}

func (p *kUserServiceClient) UpdateEmail(ctx context.Context, req *user.UpdateEmailRequest, callOptions ...callopt.Option) (r *user.UpdateEmailResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateEmail(ctx, req)
}

func (p *kUserServiceClient) UpdatePhone(ctx context.Context, req *user.UpdatePhoneRequest, callOptions ...callopt.Option) (r *user.UpdatePhoneResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePhone(ctx, req)
}

func (p *kUserServiceClient) UpdatePasswd(ctx context.Context, req *user.UpdatePasswdRequest, callOptions ...callopt.Option) (r *user.UpdatePasswdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePasswd(ctx, req)
}

func (p *kUserServiceClient) UpdateBirthDay(ctx context.Context, req *user.UpdateBirthDayRequest, callOptions ...callopt.Option) (r *user.UpdateBirthDayResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateBirthDay(ctx, req)
}
