package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/xince-fun/InstaGo/server/api/conf"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"net"
	"strconv"
)

func InitRegistry() (registry.Registry, *registry.Info) {
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		conf.GlobalConsulConf.Host,
		strconv.Itoa(conf.GlobalConsulConf.Port))
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}

	r := consul.NewConsulRegister(consulClient,
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       consts.ConsulCheckInterval,
			Timeout:                        consts.ConsulCheckTimeout,
			DeregisterCriticalServiceAfter: consts.ConsulCheckDeregisterCriticalServiceAfter,
		}))

	sf, err := snowflake.NewNode(2)
	if err != nil {
		hlog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: conf.GlobalServerConf.Name,
		Addr: utils.NewNetAddr(consts.TCP, net.JoinHostPort(conf.GlobalServerConf.Host,
			strconv.Itoa(conf.GlobalServerConf.Port))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
		Weight: registry.DefaultWeight,
	}
	return r, info
}
