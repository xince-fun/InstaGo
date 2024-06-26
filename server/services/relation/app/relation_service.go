package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/relation/domain/repo"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/cache"
	"github.com/xince-fun/InstaGo/server/services/relation/infra/sal"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/common"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"sync"
	"time"
)

var RelationApplicationSet = wire.NewSet(
	repo.RelationRepositorySet,
	cache.CacheManagerSet,
	sal.UserManagerSet,
	NewRelationApplicationService,
	wire.Bind(new(CacheManager), new(*cache.RedisManager)),
	wire.Bind(new(UserManager), new(*sal.UserManager)),
)

type RelationApplicationService struct {
	relationRepo repo.RelationRepo
	userManager  UserManager
	cacheManager CacheManager
}

type UserManager interface {
	CheckUserExist(context.Context, *user.CheckUserExistRequest) (*user.CheckUserExistResponse, error)
	GetUserInfo(context.Context, *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error)
}

type CacheManager interface {
	Get(context.Context, string, interface{}) error
	Set(context.Context, string, cache.CacheItem) error
}

func NewRelationApplicationService(relationRepo repo.RelationRepo, userManager UserManager,
	cacheManager CacheManager) *RelationApplicationService {
	return &RelationApplicationService{
		userManager:  userManager,
		relationRepo: relationRepo,
		cacheManager: cacheManager,
	}
}

func (s *RelationApplicationService) Follow(ctx context.Context, req *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
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

func (s *RelationApplicationService) Unfollow(ctx context.Context, req *relation.UnfollowRequest) (resp *relation.UnfollowResponse, err error) {
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

func (s *RelationApplicationService) CountFolloweeList(ctx context.Context, req *relation.CountFolloweeListRequest) (resp *relation.CountFolloweeListResponse, err error) {
	resp = new(relation.CountFolloweeListResponse)

	// check if user exists
	userResp, err := s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.UserId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	value := cache.RelationItem{}
	if err = s.cacheManager.Get(ctx, fmt.Sprintf(consts.FolloweeCountCacheKey, req.UserId), &value); err == nil && value.IsDirty() {
		resp.Count = value.Count
		return resp, nil
	} else {
		count, err := s.relationRepo.CountFollowee(ctx, req.UserId)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}
		resp.Count = int32(count)
		go func() {
			if err := s.cacheManager.Set(ctx, fmt.Sprintf(consts.FolloweeCountCacheKey, req.UserId), &cache.RelationItem{Count: int32(count)}); err != nil {
				klog.Infof("set followee count cache failed, err: %v", err)
			}
		}()
		return resp, nil
	}
}

func (s *RelationApplicationService) CountFollowerList(ctx context.Context, req *relation.CountFollowerListRequest) (resp *relation.CountFollowerListResponse, err error) {
	resp = new(relation.CountFollowerListResponse)

	// check if user exists
	userResp, err := s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.UserId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	value := cache.RelationItem{}
	if err = s.cacheManager.Get(ctx, fmt.Sprintf(consts.FollowerCountCacheKey, req.UserId), &value); err == nil && value.IsDirty() {
		resp.Count = value.Count
		return resp, nil
	} else {
		count, err := s.relationRepo.CountFollower(ctx, req.UserId)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}
		resp.Count = int32(count)
		go func() {
			if err := s.cacheManager.Set(ctx, fmt.Sprintf(consts.FollowerCountCacheKey, req.UserId), &cache.RelationItem{Count: int32(count)}); err != nil {
				klog.Infof("set follower count cache failed, err: %v", err)
			}
		}()
		return resp, nil
	}
}

func (s *RelationApplicationService) GetFolloweeList(ctx context.Context, req *relation.GetFolloweeListRequest) (resp *relation.GetFolloweeListResponse, err error) {
	resp = new(relation.GetFolloweeListResponse)

	// check if user exists
	userResp, err := s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.UserId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	var followeeIdList []string
	idList := cache.IDList{}
	if err = s.cacheManager.Get(ctx, fmt.Sprintf(consts.FolloweeListCacheKey, req.UserId), &idList); err == nil && idList.IsDirty() {
		for _, id := range idList {
			followeeIdList = append(followeeIdList, id)
		}
	} else {
		rList, err := s.relationRepo.GetFolloweeList(ctx, req.UserId, int(req.Offset), int(req.Limit))
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}

		c := cache.IDList{}
		for _, r := range rList {
			followeeIdList = append(followeeIdList, r.FolloweeID)
			c = append(c, r.FolloweeID)
		}
		go func() {
			if err := s.cacheManager.Set(ctx, fmt.Sprintf(consts.FolloweeListCacheKey, req.UserId), &c); err != nil {
				klog.Infof("set followee list cache failed, err: %v", err)
			}
		}()
	}
	followeeList, err := s.idListToUserList(ctx, followeeIdList)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationSrvError)
		return resp, nil
	}

	resp.FolloweeList = followeeList
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *RelationApplicationService) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	resp = new(relation.GetFollowerListResponse)

	// check if user exists
	userResp, err := s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.UserId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	var followerIdList []string
	idList := cache.IDList{}
	if err = s.cacheManager.Get(ctx, fmt.Sprintf(consts.FollowerListCacheKey, req.UserId), &idList); err == nil && idList.IsDirty() {
		for _, id := range idList {
			followerIdList = append(followerIdList, id)
		}
	} else {
		rList, err := s.relationRepo.GetFollowerList(ctx, req.UserId, int(req.Offset), int(req.Limit))
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}

		c := cache.IDList{}
		for _, r := range rList {
			followerIdList = append(followerIdList, r.FollowerID)
			c = append(c, r.FollowerID)
		}
		go func() {
			if err := s.cacheManager.Set(ctx, fmt.Sprintf(consts.FollowerListCacheKey, req.UserId), &c); err == nil && idList.IsDirty() {
				klog.Infof("set follower list cache failed, err: %v", err)
			}
		}()
	}

	followerList, err := s.idListToUserList(ctx, followerIdList)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationSrvError)
		return resp, nil
	}

	resp.FollowerList = followerList
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *RelationApplicationService) idListToUserList(ctx context.Context, idList []string) ([]*common.UserInfo, error) {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	userList := make([]*common.UserInfo, 0, len(idList))
	maxRetries := 3
	retryInternal := 200 * time.Millisecond

	for _, id := range idList {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			retryCount := 0
			for retryCount < maxRetries {
				resp, err := s.userManager.GetUserInfo(ctx, &user.GetUserInfoRequest{UserId: id})
				if err != nil {
					retryCount++
					klog.Errorf("get user info %s failed %d times, err: %v", id, retryCount, err)
					time.Sleep(retryInternal)
					continue
				} else {
					mu.Lock()
					userList = append(userList, &common.UserInfo{
						UserId:   resp.UserInfo.UserId,
						Account:  resp.UserInfo.Account,
						FullName: resp.UserInfo.FullName,
						Avatar:   resp.UserInfo.Avatar,
					})
					mu.Unlock()
					break
				}
			}

			if retryCount >= maxRetries {
				klog.Errorf("get user info failed, id: %s", id)
			}
		}(id)
	}

	wg.Wait()

	return userList, nil
}
