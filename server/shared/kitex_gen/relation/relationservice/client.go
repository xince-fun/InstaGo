// Code generated by Kitex v0.9.1. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Follow(ctx context.Context, req *relation.FollowRequest, callOptions ...callopt.Option) (r *relation.FollowResponse, err error)
	Unfollow(ctx context.Context, req *relation.UnfollowRequest, callOptions ...callopt.Option) (r *relation.UnfollowResponse, err error)
	CountFolloweeList(ctx context.Context, req *relation.CountFolloweeListRequest, callOptions ...callopt.Option) (r *relation.CountFolloweeListResponse, err error)
	CountFollowerList(ctx context.Context, req *relation.CountFollowerListRequest, callOptions ...callopt.Option) (r *relation.CountFollowerListResponse, err error)
	GetFolloweeList(ctx context.Context, req *relation.GetFolloweeListRequest, callOptions ...callopt.Option) (r *relation.GetFolloweeListResponse, err error)
	GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest, callOptions ...callopt.Option) (r *relation.GetFollowerListResponse, err error)
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
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) Follow(ctx context.Context, req *relation.FollowRequest, callOptions ...callopt.Option) (r *relation.FollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, req)
}

func (p *kRelationServiceClient) Unfollow(ctx context.Context, req *relation.UnfollowRequest, callOptions ...callopt.Option) (r *relation.UnfollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Unfollow(ctx, req)
}

func (p *kRelationServiceClient) CountFolloweeList(ctx context.Context, req *relation.CountFolloweeListRequest, callOptions ...callopt.Option) (r *relation.CountFolloweeListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFolloweeList(ctx, req)
}

func (p *kRelationServiceClient) CountFollowerList(ctx context.Context, req *relation.CountFollowerListRequest, callOptions ...callopt.Option) (r *relation.CountFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFollowerList(ctx, req)
}

func (p *kRelationServiceClient) GetFolloweeList(ctx context.Context, req *relation.GetFolloweeListRequest, callOptions ...callopt.Option) (r *relation.GetFolloweeListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFolloweeList(ctx, req)
}

func (p *kRelationServiceClient) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest, callOptions ...callopt.Option) (r *relation.GetFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, req)
}
