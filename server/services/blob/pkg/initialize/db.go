package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xince-fun/InstaGo/server/services/blob/conf"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
	"time"
)

func InitDB() *gorm.DB {
	c := conf.GlobalServerConf.MysqlConfig
	dsn := fmt.Sprintf(consts.MysqlDSN, c.User, c.Password, c.Host, c.Port, c.DB)

	l := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Minute,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: l,
		})

	if err != nil {
		klog.Fatalf("database init mysql gorm open failed: %s", err.Error())
	}

	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}

	if err = db.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		klog.Fatalf("sqlDB open error: %s", err.Error())
	}
	db = db.Debug()

	sqlDB.SetConnMaxIdleTime(time.Duration(conf.GlobalServerConf.MysqlConfig.MaxIdleConns) * time.Minute)
	sqlDB.SetMaxIdleConns(conf.GlobalServerConf.MysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.GlobalServerConf.MysqlConfig.MaxOpenConns)

	return db
}
