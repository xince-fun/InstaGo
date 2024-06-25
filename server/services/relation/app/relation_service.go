package app

import (
	"context"
	"errors"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/repo"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/utils"
)

type FollowApplicationService struct {
	relationRepo repo.RelationRepo
	userManager  UserManager
}

type UserManager interface {
	CheckUserExist(context.Context, *user.CheckUserExistRequest) (*user.CheckUserExistResponse, error)
}

func NewFollowApplicationService(relationRepo repo.RelationRepo, userManager UserManager) *FollowApplicationService {
	return &FollowApplicationService{
		userManager:  userManager,
		relationRepo: relationRepo,
	}
}

func (s *FollowApplicationService) Follow(ctx context.Context, req *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	resp = new(relation.FollowResponse)

	// check if follower exists
	userResp, err := s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.FollowerId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	if req.FolloweeId == req.FollowerId {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationSelfError)
		return resp, nil
	}

	// check followee exists
	userResp, err = s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.FolloweeId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	r := entity.Relation{
		FollowerID: req.FollowerId,
		FolloweeID: req.FolloweeId,
	}

	if err = s.relationRepo.UpsertRelation(ctx, &r); errors.Is(err, errno.RelationDBError) {
		resp.BaseResp = utils.BuildBaseResp(err)
		return resp, nil
	} else if errors.Is(err, errno.RelationExistError) {
		resp.BaseResp = utils.BuildBaseResp(err)
		return resp, nil
	} else if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *FollowApplicationService) Unfollow(ctx context.Context, req *relation.UnfollowRequest) (resp *relation.UnfollowResponse, err error) {
	resp = new(relation.UnfollowResponse)

	// check if follower exists
	userResp, err := s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.FollowerId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	if req.FolloweeId == req.FollowerId {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationSelfError)
		return resp, nil
	}

	// check followee exists
	userResp, err = s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.FolloweeId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	r := entity.Relation{
		FollowerID: req.FollowerId,
		FolloweeID: req.FolloweeId,
	}

	if err = s.relationRepo.DeleteRelation(ctx, &r); errors.Is(err, errno.RelationDBError) {
		resp.BaseResp = utils.BuildBaseResp(err)
		return resp, nil
	} else if errors.Is(err, errno.RelationNotExistError) {
		resp.BaseResp = utils.BuildBaseResp(err)
		return resp, nil
	} else if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}
