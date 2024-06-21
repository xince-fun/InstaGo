// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xince-fun/InstaGo/server/api/pkg/initialize/rpc"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/xince-fun/InstaGo/server/api/biz/model/user"
	kuser "github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
)

// Login .
// @router /api/v1/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest

	resp := new(kuser.LoginResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	phoneOrEmailType, err := PhoneOrEmail(req.Account)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr.WithMessage(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	switch phoneOrEmailType {
	case consts.PhoneNumber:
		resp, err = rpc.LoginPhone(ctx, &kuser.LoginPhoneRequest{
			Account: req.Account,
			Passwd:  req.Passwd,
		})
		if err != nil {
			hlog.Error("rpc user service err", err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	case consts.Email:
		resp, err = rpc.LoginEmail(ctx, &kuser.LoginEmailRequest{
			Account: req.Account,
			Passwd:  req.Passwd,
		})
		if err != nil {
			hlog.Error("rpc user service err", err)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// Register .
// @router /api/v1/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest

	resp := new(kuser.RegisterResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	phoneOrEmailType, err := PhoneOrEmail(req.PhoneOrEmail)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr.WithMessage(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	switch phoneOrEmailType {
	case consts.PhoneNumber:
		resp, err = rpc.RegisterPhone(ctx, &kuser.RegisterPhoneRequest{
			PhoneNumber: req.PhoneOrEmail,
			Passwd:      req.Passwd,
			Account:     req.Account,
			FullName:    req.FullName,
		})
		if err != nil {
			hlog.Error("rpc user service err", err)
			resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	case consts.Email:
		resp, err = rpc.RegisterEmail(ctx, &kuser.RegisterEmailRequest{
			Email:    req.PhoneOrEmail,
			Passwd:   req.Passwd,
			Account:  req.Account,
			FullName: req.FullName,
		})
		if err != nil {
			hlog.Error("rpc user service err", err)
			resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func PhoneOrEmail(phoneOrEmail string) (int, error) {
	if utils.IsValidRegexp(consts.PhoneNumberRegexp, phoneOrEmail) {
		return consts.PhoneNumber, nil
	}
	if utils.IsValidRegexp(consts.EmailRegexp, phoneOrEmail) {
		return consts.Email, nil
	}
	return consts.Unknown, errno.ParamsErr.WithMessage("phone or email format error")
}

// UpdateEmail .
// @router /api/v1/user/update_email [POST]
func UpdateEmail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateEmailRequest

	resp := new(kuser.UpdateEmailResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = rpc.UpdateEmail(ctx, &kuser.UpdateEmailRequest{
		UserId: c.MustGet(consts.UserID).(string),
		Email:  req.Email,
	})
	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdatePhone .
// @router /api/v1/user/update_phone [POST]
func UpdatePhone(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdatePhoneRequest
	resp := new(kuser.UpdatePhoneResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = rpc.UpdatePhone(ctx, &kuser.UpdatePhoneRequest{
		UserId:      c.MustGet(consts.UserID).(string),
		PhoneNumber: req.Phone,
	})
	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdatePasswd .
// @router /api/v1/user/update_passwd [POST]
func UpdatePasswd(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdatePasswdRequest

	resp := new(kuser.UpdatePasswdResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = rpc.UpdatePasswd(ctx, &kuser.UpdatePasswdRequest{
		UserId:     c.MustGet(consts.UserID).(string),
		OldPasswd:  req.OldPasswd,
		NewPasswd_: req.NewPasswd,
	})
	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateBirthDay .
// @router /api/v1/user/update_birthday [POST]
func UpdateBirthDay(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateBirthDayRequest

	resp := new(kuser.UpdateBirthDayResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = rpc.UpdateBirthDay(ctx, &kuser.UpdateBirthDayRequest{
		UserId: c.MustGet(consts.UserID).(string),
		Year:   req.Year,
		Month:  req.Month,
		Day:    req.Day,
	})
	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = utils.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
