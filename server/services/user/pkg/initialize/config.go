package initialize

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"github.com/xince-fun/InstaGo/server/shared/utils"

	"github.com/xince-fun/InstaGo/server/services/user/conf"
	"net"
	"path/filepath"
	"strconv"
)

const (
	PREFIX    = "conf"
	CONF_TEST = "conf-test.yaml"
	CONF_PROD = "conf-prod.yaml"
	CONF_DEV  = "conf-dev.yaml"
)

func InitConfig() {
	v := viper.New()
	confFilePelPath := filepath.Join(PREFIX, CONF_DEV)
	v.SetConfigFile(confFilePelPath)
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed")
	}
	if err := v.Unmarshal(&conf.GlobalConsulConf); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err.Error())
	}

	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		conf.GlobalConsulConf.Host,
		strconv.Itoa(conf.GlobalConsulConf.Port),
	)
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}
	content, _, err := consulClient.KV().Get(conf.GlobalConsulConf.Key, nil)
	if err != nil {
		klog.Fatalf("consul kv failed: %s", err.Error())
	}
	err = sonic.Unmarshal(content.Value, &conf.GlobalServerConf)
	if err != nil {
		klog.Fatalf("sonic unmarshal conf failed: %s", err.Error())
	}

	if conf.GlobalServerConf.Host == "" {
		conf.GlobalServerConf.Host, err = utils.GetLocalIPv4Address()
		if err != nil {
			klog.Fatalf("get localIpv4Addr failed: %s", err.Error())
		}
	}
}
