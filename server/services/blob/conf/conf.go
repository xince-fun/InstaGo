package conf

var (
	GlobalServerConf *ServerConfig
	GlobalConsulConf *ConsulConfig
)

type DBConfig struct {
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	DB              string `mapstructure:"db" json:"db"`
	User            string `mapstructure:"user" json:"user"`
	Password        string `mapstructure:"password" json:"password"`
	Salt            string `mapstructure:"salt" json:"salt"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" json:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime"`
}

type ServerConfig struct {
	Env          string       `mapstructure:"env" json:"env"`
	Name         string       `mapstructure:"name" json:"name"`
	Host         string       `mapstructure:"host" json:"host"`
	OtelConfig   OtelConfig   `mapstructure:"otel" json:"otel"`
	DBConfig     DBConfig     `mapstructure:"db" json:"db"`
	BucketConfig BucketConfig `mapstructure:"bucket" json:"bucket"`
	MQConfig     MQConfig     `mapstructure:"mq" json:"mq"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type BucketConfig struct {
	EndPoint     string `mapstructure:"endpoint" json:"endpoint"`
	AccessKeyID  string `mapstructure:"access_key_id" json:"access_key_id"`
	AccessSecret string `mapstructure:"access_secret" json:"access_secret"`
	AvatarBucket string `mapstructure:"avatar_bucket" json:"avatar_bucket"`
}

type MQConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Exchange string `mapstructure:"exchange" json:"exchange"`
	User     string `mapstructure:"user" json:"user"`
	Passwd   string `mapstructure:"passwd" json:"passwd"`
}
