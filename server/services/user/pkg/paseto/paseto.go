package paseto

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/hertz-contrib/paseto"
	"github.com/xince-fun/InstaGo/server/services/user/conf"
)

var TokenGeneratorSet = wire.NewSet(
	ProvidePasetoAsymmetricKey,
	ProvidePasetoImplicit,
	NewTokenGenerator,
)

func ProvidePasetoAsymmetricKey() AsymmetricKey {
	return AsymmetricKey(conf.GlobalServerConf.PasetoConfig.SecretKey)
}

func ProvidePasetoImplicit() Implicit {
	return Implicit(conf.GlobalServerConf.PasetoConfig.Implicit)
}

type AsymmetricKey string
type Implicit []byte

type TokenGenerator struct {
	paseto.GenTokenFunc
}

func NewTokenGenerator(asymmetricKey AsymmetricKey, implicit Implicit) *TokenGenerator {
	signFunc, err := paseto.NewV4SignFunc(string(asymmetricKey), implicit)
	if err != nil {
		klog.Fatalf("paseto token generator failed: %s", err.Error())
		return nil
	}
	return &TokenGenerator{signFunc}
}

func (g *TokenGenerator) CreateToken(claims *paseto.StandardClaims) (token string, err error) {
	return g.GenTokenFunc(claims, nil, nil)
}
