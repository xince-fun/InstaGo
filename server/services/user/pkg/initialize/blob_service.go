package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/xince-fun/InstaGo/server/services/user/conf"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob/blobservice"
)

func InitBlob() blobservice.Client {
	r, err := consul.NewConsulResolver(fmt.Sprintf("%s:%d",
		conf.GlobalConsulConf.Host,
		conf.GlobalConsulConf.Port))
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(conf.GlobalServerConf.BlobSrvInfo.Name),
		provider.WithExportEndpoint(conf.GlobalServerConf.OtelConfig.EndPoint),
		provider.WithInsecure(),
	)

	c, err := blobservice.NewClient(
		conf.GlobalServerConf.BlobSrvInfo.Name,
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GlobalServerConf.BlobSrvInfo.Name}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	return c
}
