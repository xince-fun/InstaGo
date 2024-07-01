package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/xince-fun/InstaGo/server/api/conf"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	krelation "github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation"
	relation "github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation/relationservice"
	"github.com/xince-fun/InstaGo/server/shared/utils"
)

var relationClient relation.Client

func initRelation() {
	r, err := consul.NewConsulResolver(fmt.Sprintf("%s:%d",
		conf.GlobalConsulConf.Host,
		conf.GlobalConsulConf.Port))
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(conf.GlobalServerConf.UserSrvInfo.Name),
		provider.WithExportEndpoint(conf.GlobalServerConf.OtelConfig.EndPoint),
		provider.WithInsecure(),
	)

	// create a new client
	c, err := relation.NewClient(
		conf.GlobalServerConf.RelationSrvInfo.Name,
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GlobalServerConf.RelationSrvInfo.Name}),
	)

	if err != nil {
		klog.Fatalf("cannot init client: %v\n", err)
	}

	relationClient = c
}

func Follow(ctx context.Context, req *krelation.FollowRequest) (resp *krelation.FollowResponse, err error) {
	resp, err = relationClient.Follow(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func Unfollow(ctx context.Context, req *krelation.UnfollowRequest) (resp *krelation.UnfollowResponse, err error) {
	resp, err = relationClient.Unfollow(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func CountFollowee(ctx context.Context, req *krelation.CountFolloweeListRequest) (resp *krelation.CountFolloweeListResponse, err error) {
	resp, err = relationClient.CountFolloweeList(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func CountFollower(ctx context.Context, req *krelation.CountFollowerListRequest) (resp *krelation.CountFollowerListResponse, err error) {
	resp, err = relationClient.CountFollowerList(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func GetFolloweeList(ctx context.Context, req *krelation.GetFolloweeListRequest) (resp *krelation.GetFolloweeListResponse, err error) {
	resp, err = relationClient.GetFolloweeList(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func GetFollowerList(ctx context.Context, req *krelation.GetFollowerListRequest) (resp *krelation.GetFollowerListResponse, err error) {
	resp, err = relationClient.GetFollowerList(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func IsFollow(ctx context.Context, req *krelation.IsFollowRequest) (resp *krelation.IsFollowResponse, err error) {
	resp, err = relationClient.IsFollow(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}
