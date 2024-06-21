package common

import (
	pt "aidanwoods.dev/go-paseto"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/paseto"
	"github.com/xince-fun/InstaGo/server/api/conf"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/middleware"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"net/http"
)

func CommonMW() []app.HandlerFunc {
	return []app.HandlerFunc{
		// use cors mw
		middleware.Cors(),
		// use recovery mw
		middleware.Recovery(),
		// use gzip mw
		gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".jpg", ".mp4", ".png"})),
	}
}

func PasetoAuth(audience string) app.HandlerFunc {
	c := conf.GlobalServerConf.PasetoInfo
	pf, err := paseto.NewV4PublicParseFunc(c.PubKey, []byte(c.Implicit), paseto.WithAudience(audience), paseto.WithNotBefore())
	if err != nil {
		hlog.Fatal(err)
	}
	sh := func(ctx context.Context, c *app.RequestContext, token *pt.Token) {
		aid, err := token.GetString("id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.BuildBaseResp(errno.BadRequest.WithMessage("missing userID in token")))
		}
		c.Set(consts.UserID, aid)
	}

	eh := func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusUnauthorized, utils.BuildBaseResp(errno.BadRequest.WithMessage("invalid token")))
		c.Abort()
	}
	return paseto.New(paseto.WithTokenPrefix("Bearer "), paseto.WithParseFunc(pf), paseto.WithSuccessHandler(sh), paseto.WithErrorFunc(eh))
}
