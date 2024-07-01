package main

import (
	"context"
	"github.com/xince-fun/InstaGo/server/services/post/app"
	post "github.com/xince-fun/InstaGo/server/shared/kitex_gen/post"
)

// PostServiceImpl implements the last service interface defined in the IDL.
type PostServiceImpl struct {
	app *app.PostApplicationService
}

// PostPhoto implements the PostServiceImpl interface.
func (s *PostServiceImpl) PostPhoto(ctx context.Context, req *post.PostPhotoRequest) (resp *post.PostPhotoResponse, err error) {
	return s.app.PostPhoto(ctx, req)
}
