package app

import (
	"context"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/post"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"time"
)

type PostApplicationService struct {
	userManager UserManager
	blobManager BlobManager
	detector    Detector
}

type UserManager interface {
	CheckUserExist(context.Context, *user.CheckUserExistRequest) (*user.CheckUserExistResponse, error)
}

type BlobManager interface {
	UploadBlob(context.Context, *blob.GeneratePutPreSignedUrlRequest) (*blob.GeneratePutPreSignedUrlResponse, error)
	GetBlob(context.Context, *blob.GenerateGetPreSignedUrlRequest) (*blob.GenerateGetPreSignedUrlResponse, error)
	NotifyBlobUpload(context.Context, *blob.NotifyBlobUploadRequest) (*blob.NotifyBlobUploadResponse, error)
}

type Detector interface {
	DetectPhoto([]byte) error
	DetectVideo([]byte) error
}

func NewPostApplicationService(userManager UserManager, blobManager BlobManager, detector Detector) *PostApplicationService {
	return &PostApplicationService{
		userManager: userManager,
		blobManager: blobManager,
		detector:    detector,
	}
}

// PostPhoto 调用blob服务生成预签名URL,返回给客户端
func (s *PostApplicationService) PostPhoto(ctx context.Context, req *post.PostPhotoRequest) (resp *post.PostPhotoResponse, err error) {
	resp = new(post.PostPhotoResponse)

	userResp, err := s.userManager.CheckUserExist(ctx, &user.CheckUserExistRequest{UserId: req.UserId})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if !userResp.IsExist {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}

	n := len(req.Photo)
	if n > consts.DetectorBufferSize {
		n = consts.DetectorBufferSize
	}
	buffer := make([]byte, consts.DetectorBufferSize)
	copy(buffer, req.Photo[:n])

	if err = s.detector.DetectPhoto(buffer); err != nil {
		resp.BaseResp = utils.BuildBaseResp(err)
		return resp, nil
	}

	blobResp, err := s.blobManager.UploadBlob(ctx, &blob.GeneratePutPreSignedUrlRequest{
		UserId:   req.UserId,
		BlobType: consts.PhotoBlobType,
		Timeout:  int32(10 * time.Second.Seconds()),
	})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	resp.PhotoUrl = blobResp.Url
	resp.PostId = blobResp.Id
	resp.ObjectName = blobResp.ObjectName
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}
