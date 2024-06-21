package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/xince-fun/InstaGo/server/services/blob/conf"
	"github.com/xince-fun/InstaGo/server/shared/consts"
)

func InitMQ() *amqp.Connection {
	c := conf.GlobalServerConf.MQConfig
	amqpConn, err := amqp.Dial(fmt.Sprintf(consts.AmqpURI, c.User, c.Passwd, c.Host, c.Port))
	if err != nil {
		klog.Fatalf("init mq failed: %v", err)
	}
	return amqpConn
}
