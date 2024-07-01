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
	scache "github.com/xince-fun/InstaGo/server/shared/cache"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/common"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"strconv"
	"sync"
	"time"
)

var RelationApplicationSet = wire.NewSet(
	repo.RelationRepositorySet,
	cache.CacheManagerSet,
	sal.UserManagerSet,
	NewRelationApplicationService,
	wire.Bind(new(scache.CacheManager), new(*cache.CacheManager)),
	wire.Bind(new(UserManager), new(*sal.UserManager)),
)

type RelationApplicationService struct {
	relationRepo repo.RelationRepo
	userManager  UserManager
	cacheManager scache.CacheManager
}

type UserManager interface {
	CheckUserExist(context.Context, *user.CheckUserExistRequest) (*user.CheckUserExistResponse, error)
	GetUserInfo(context.Context, *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error)
}

func NewRelationApplicationService(relationRepo repo.RelationRepo, userManager UserManager,
	cacheManager scache.CacheManager) *RelationApplicationService {
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

	if err = s.updateFolloweeListCache(ctx, req.FollowerId, r, true); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationCacheError)
		return resp, nil
	}

	if err = s.updateFollowerListCache(ctx, req.FolloweeId, r, true); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationCacheError)
		return resp, nil
	}

	if err = s.updateFolloweeCountCache(ctx, req.FollowerId, true); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationCacheError)
		return resp, nil
	}

	if err = s.updateFollowerCountCache(ctx, req.FolloweeId, true); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RelationCacheError)
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

	if countString, err := s.cacheManager.Get(ctx, fmt.Sprintf(consts.FolloweeCountCacheKey, req.UserId)); err == nil && countString != "" {
		count, err := strconv.ParseUint(countString, 10, 32)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationCacheError)
			return resp, nil
		}
		resp.Count = int32(count)
		return resp, nil
	} else {
		count, err := s.relationRepo.CountFollowee(ctx, req.UserId)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}
		resp.Count = int32(count)
		go func() {
			countString = strconv.FormatUint(uint64(count), 10)
			if err := s.cacheManager.Set(ctx, fmt.Sprintf(consts.FolloweeCountCacheKey, req.UserId), countString); err != nil {
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

	if countString, err := s.cacheManager.Get(ctx, fmt.Sprintf(consts.FollowerCountCacheKey, req.UserId)); err == nil && countString != "" {
		count, err := strconv.ParseUint(countString, 10, 32)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationCacheError)
			return resp, nil
		}
		resp.Count = int32(count)
		return resp, nil
	} else {
		count, err := s.relationRepo.CountFollower(ctx, req.UserId)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}
		resp.Count = int32(count)
		go func() {
			countString = strconv.FormatUint(uint64(count), 10)
			if err := s.cacheManager.Set(ctx, fmt.Sprintf(consts.FollowerCountCacheKey, req.UserId), countString); err != nil {
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

	cacheKey := fmt.Sprintf(consts.FolloweeListCacheKey, req.UserId)
	idList, err := s.cacheManager.Client().SMembers(ctx, cacheKey).Result()

	followeeIdList := make([]string, len(idList))

	if len(idList) == 0 || err != nil {
		rList, err := s.relationRepo.GetFolloweeList(ctx, req.UserId, int(req.Offset), int(req.Limit))
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}

		followeeIdList = make([]string, len(rList))
		for i, r := range rList {
			followeeIdList[i] = r.FolloweeID
			s.cacheManager.Client().SAdd(ctx, cacheKey, r.FolloweeID)
		}
	} else {
		for i, id := range idList {
			followeeIdList[i] = id
		}
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

	cacheKey := fmt.Sprintf(consts.FollowerListCacheKey, req.UserId)
	idList, err := s.cacheManager.Client().SMembers(ctx, cacheKey).Result()
	followerIdList := make([]string, len(idList))

	if len(idList) == 0 || err != nil {
		rList, err := s.relationRepo.GetFollowerList(ctx, req.UserId, int(req.Offset), int(req.Limit))
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}

		followerIdList = make([]string, len(rList))
		for i, r := range rList {
			followerIdList[i] = r.FollowerID
			s.cacheManager.Client().SAdd(ctx, cacheKey, r.FollowerID)
		}
	} else {
		for i, id := range idList {
			followerIdList[i] = id
		}
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

func (s *RelationApplicationService) IsFollow(ctx context.Context, req *relation.IsFollowRequest) (resp *relation.IsFollowResponse, err error) {
	resp = new(relation.IsFollowResponse)

	cacheKey := fmt.Sprintf(consts.IsFollowCacheKey, req.FollowerId, req.FolloweeId)
	if isFollow, err := s.cacheManager.Get(ctx, cacheKey); err == nil && isFollow != "" {
		resp.IsFollow = isFollow == "true"
		return resp, nil
	} else {
		res, err := s.relationRepo.IsFollow(ctx, req.FollowerId, req.FolloweeId)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.RelationDBError)
			return resp, nil
		}
		resp.IsFollow = res
		resp.BaseResp = utils.BuildBaseResp(nil)

		// write back to cache
		go func() {
			if err := s.cacheManager.Set(ctx, cacheKey, strconv.FormatBool(res)); err != nil {
				klog.Infof("write back to cache error: %v", err)
			}
		}()
		return resp, nil
	}
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

func (s *RelationApplicationService) updateFolloweeListCache(ctx context.Context, userId string, relation entity.Relation, op bool) error {
	cacheKey := fmt.Sprintf(consts.FolloweeListCacheKey, userId)
	if op {
		if err := s.cacheManager.Client().SAdd(ctx, cacheKey, relation.FolloweeID).Err(); err != nil {
			klog.Errorf("add followee list cache failed, err: %v", err)
			return err
		}
	} else {
		if err := s.cacheManager.Client().SRem(ctx, cacheKey, relation.FolloweeID).Err(); err != nil {
			klog.Errorf("remove followee list cache failed, err: %v", err)
			return err
		}
	}
	return nil
}

func (s *RelationApplicationService) updateFollowerListCache(ctx context.Context, userId string, relation entity.Relation, op bool) error {
	cacheKey := fmt.Sprintf(consts.FollowerListCacheKey, userId)
	if op {
		if err := s.cacheManager.Client().SAdd(ctx, cacheKey, relation.FollowerID).Err(); err != nil {
			klog.Errorf("add follower list cache failed, err: %v", err)
			return err
		}
	} else {
		if err := s.cacheManager.Client().SRem(ctx, cacheKey, relation.FollowerID).Err(); err != nil {
			klog.Errorf("remove follower list cache failed, err: %v", err)
			return err
		}
	}
	return nil
}

func (s *RelationApplicationService) updateFolloweeCountCache(ctx context.Context, userId string, op bool) error {
	cacheKey := fmt.Sprintf(consts.FolloweeCountCacheKey, userId)
	var count uint32

	oriCountString, err := s.cacheManager.Get(ctx, cacheKey)
	if err != nil {
		if !(errors.Is(err, cache.ErrCacheNotFound) || errors.Is(err, cache.ErrCacheNotFound)) {
			return err
		}
	}
	// not hit
	if errors.Is(err, cache.ErrCacheNotFound) {
		oriCount, err := s.relationRepo.CountFollowee(ctx, userId)
		if err != nil {
			return err
		}
		if op {
			oriCount = oriCount + 1
		} else {
			if oriCount == 0 {
				oriCount = 0
			} else {
				oriCount = oriCount - 1
			}
		}

		count = uint32(oriCount)
	} else {
		oriCount, err := strconv.ParseUint(oriCountString, 10, 32)
		if err != nil {
			return err
		}
		oriCountInt := uint32(oriCount)
		if op {
			count = oriCountInt + 1
		} else {
			if oriCountInt == 0 {
				count = 0
			} else {
				count = oriCountInt - 1
			}
		}
	}

	// write back
	countString := strconv.FormatUint(uint64(count), 10)
	return s.cacheManager.Set(ctx, cacheKey, countString)
}

func (s *RelationApplicationService) updateFollowerCountCache(ctx context.Context, userId string, op bool) error {
	cacheKey := fmt.Sprintf(consts.FollowerListCacheKey, userId)
	var count uint32

	oriCountString, err := s.cacheManager.Get(ctx, cacheKey)
	if err != nil {
		if !(errors.Is(err, cache.ErrCacheNotFound) || errors.Is(err, cache.ErrCacheNotFound)) {
			return err
		}
	}
	// not hit
	if errors.Is(err, cache.ErrCacheNotFound) {
		oriCount, err := s.relationRepo.CountFollower(ctx, userId)
		if err != nil {
			return err
		}
		if op {
			oriCount = oriCount + 1
		} else {
			if oriCount == 0 {
				oriCount = 0
			} else {
				oriCount = oriCount - 1
			}
		}

		count = uint32(oriCount)
	} else {
		oriCount, err := strconv.ParseUint(oriCountString, 10, 32)
		if err != nil {
			return err
		}
		oriCountInt := uint32(oriCount)
		if op {
			count = oriCountInt + 1
		} else {
			if oriCountInt == 0 {
				count = 0
			} else {
				count = oriCountInt - 1
			}
		}
	}

	// write back
	countString := strconv.FormatUint(uint64(count), 10)
	return s.cacheManager.Set(ctx, cacheKey, countString)
}
