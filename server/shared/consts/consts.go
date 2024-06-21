package consts

import "time"

const (
	Issuer = "InstaGo"
	User   = "User"

	FiftyDays = time.Hour * 24 * 30

	UserID = "userID"

	MysqlDSN    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	AmqpURI     = "amqp://%s:%s@%s:%d"
	CorsAddress = "http://localhost:3000"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	IPFlagName    = "ip"
	IPFlagValue   = "0.0.0.0"
	IPFlagUsage   = "address"
	PortFlagName  = "port"
	PortFlagUsage = "port"

	TCP             = "tcp"
	FreePortAddress = "localhost:0"

	KlogFilePath = "./tmp/klog/logs/"
	HlogFilePath = "./tmp/hlog/logs/"

	PhoneNumberRegexp = `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
	EmailRegexp       = "\\w+@\\w+\\.[a-z]+(\\.[a-z]+)?"

	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

const (
	Unknown = iota
	PhoneNumber
	Email
)
