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
	PasetoConfig PasetoConfig `mapstructure:"paseto" json:"paseto"`
	DBConfig     DBConfig     `mapstructure:"db" json:"db"`
	RedisConfig  RedisConfig  `mapstructure:"redis" json:"redis"`
	BlobSrvInfo  RPCSrvConfig `mapstructure:"blob_srv" json:"blob_srv"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type PasetoConfig struct {
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	Implicit  string `mapstructure:"implicit" json:"implicit"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type RedisConfig struct {
	RedisServerConfig []RedisServerConfig `mapstructure:"server" json:"server"`
	LocalCacheTime    int                 `mapstructure:"local_cache" json:"local_cache"`
}

type RedisServerConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Addr string `mapstructure:"addr" json:"addr"`
}
