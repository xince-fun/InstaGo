package detector

import (
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"net/http"
)

type PostDetector struct {
}

func (d *PostDetector) DetectPhoto(data []byte) error {
	contentType := http.DetectContentType(data)
	if contentType != consts.PNGContentType && contentType != consts.JPEGContentType {
		return errno.InvalidPhotoError
	}
	return nil
}

func (d *PostDetector) DetectVideo(data []byte) error {
	contentType := http.DetectContentType(data)
	if contentType != consts.MP4ContentType {
		return errno.InvalidVideoError
	}
	return nil
}
