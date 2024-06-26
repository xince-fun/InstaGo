package main

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/relation/app"
	relation "github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation"
)

var RelationServiceImplSet = wire.NewSet(
	app.RelationApplicationSet,
	NewRelationServiceImpl,
)

func NewRelationServiceImpl(relationAppService *app.RelationApplicationService) *RelationServiceImpl {
	return &RelationServiceImpl{
		app: relationAppService,
	}
}

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct {
	app *app.RelationApplicationService
}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	return s.app.Follow(ctx, req)
}

// Unfollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Unfollow(ctx context.Context, req *relation.UnfollowRequest) (resp *relation.UnfollowResponse, err error) {
	return s.app.Unfollow(ctx, req)
}

// CountFolloweeList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFolloweeList(ctx context.Context, req *relation.CountFolloweeListRequest) (resp *relation.CountFolloweeListResponse, err error) {
	return s.app.CountFolloweeList(ctx, req)
}

// CountFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollowerList(ctx context.Context, req *relation.CountFollowerListRequest) (resp *relation.CountFollowerListResponse, err error) {
	return s.app.CountFollowerList(ctx, req)
}

// GetFolloweeList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFolloweeList(ctx context.Context, req *relation.GetFolloweeListRequest) (resp *relation.GetFolloweeListResponse, err error) {
	return s.app.GetFolloweeList(ctx, req)
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	return s.app.GetFollowerList(ctx, req)
}
