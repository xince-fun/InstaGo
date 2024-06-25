package main

import (
	"context"
	relation "github.com/xince-fun/InstaGo/server/shared/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	// TODO: Your code here...
	return
}

// Unfollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Unfollow(ctx context.Context, req *relation.UnfollowRequest) (resp *relation.UnfollowResponse, err error) {
	// TODO: Your code here...
	return
}
