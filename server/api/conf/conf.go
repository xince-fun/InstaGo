package conf

var (
	GlobalConsulConf *ConsulConfig
	GlobalServerConf *ServerConfig
)

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type PasetoConfig struct {
	PubKey   string `mapstructure:"pub_key" json:"pub_key"`
	Implicit string `mapstructure:"implicit" json:"implicit"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name            string       `mapstructure:"name" json:"name"`
	Host            string       `mapstructure:"host" json:"host"`
	Port            int          `mapstructure:"port" json:"port"`
	OtelConfig      OtelConfig   `mapstructure:"otel" json:"otel"`
	PasetoInfo      PasetoConfig `mapstructure:"paseto" json:"paseto"`
	UserSrvInfo     RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	RelationSrvInfo RPCSrvConfig `mapstructure:"relation_srv" json:"relation_srv"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
