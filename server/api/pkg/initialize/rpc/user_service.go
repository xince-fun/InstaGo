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
	kuser "github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/utils"

	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user/userservice"
)

var userClient userservice.Client

func initUser() {
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
	c, err := userservice.NewClient(
		conf.GlobalServerConf.UserSrvInfo.Name,
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GlobalServerConf.UserSrvInfo.Name}),
	)

	if err != nil {
		klog.Fatalf("cannot init client: %v\n", err)
	}

	userClient = c
}

func RegisterPhone(ctx context.Context, req *kuser.RegisterPhoneRequest) (resp *kuser.RegisterResponse, err error) {
	resp, err = userClient.RegisterPhone(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func RegisterEmail(ctx context.Context, req *kuser.RegisterEmailRequest) (resp *kuser.RegisterResponse, err error) {
	resp, err = userClient.RegisterEmail(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func LoginPhone(ctx context.Context, req *kuser.LoginPhoneRequest) (resp *kuser.LoginResponse, err error) {
	resp, err = userClient.LoginPhone(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func LoginEmail(ctx context.Context, req *kuser.LoginEmailRequest) (resp *kuser.LoginResponse, err error) {
	resp, err = userClient.LoginEmail(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func UpdateEmail(ctx context.Context, req *kuser.UpdateEmailRequest) (resp *kuser.UpdateEmailResponse, err error) {
	resp, err = userClient.UpdateEmail(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func UpdatePhone(ctx context.Context, req *kuser.UpdatePhoneRequest) (resp *kuser.UpdatePhoneResponse, err error) {
	resp, err = userClient.UpdatePhone(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func UpdatePasswd(ctx context.Context, req *kuser.UpdatePasswdRequest) (resp *kuser.UpdatePasswdResponse, err error) {
	resp, err = userClient.UpdatePasswd(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}

func UpdateBirthDay(ctx context.Context, req *kuser.UpdateBirthDayRequest) (resp *kuser.UpdateBirthDayResponse, err error) {
	resp, err = userClient.UpdateBirthDay(ctx, req)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		return resp, err
	}
	return resp, nil
}
