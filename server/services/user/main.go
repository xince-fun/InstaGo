package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/xince-fun/InstaGo/server/services/user/conf"
	"github.com/xince-fun/InstaGo/server/services/user/pkg/initialize"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	user "github.com/xince-fun/InstaGo/server/shared/kitex_gen/user/userservice"
	"net"
	"strconv"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitRegistry(Port)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(conf.GlobalServerConf.Name),
		provider.WithExportEndpoint(conf.GlobalServerConf.OtelConfig.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	svr := user.NewServer(InitializeService(),
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GlobalServerConf.Name}),
	)

	err := svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
