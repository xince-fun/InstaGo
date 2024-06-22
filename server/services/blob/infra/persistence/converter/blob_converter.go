package converter

import (
	"github.com/xince-fun/InstaGo/server/services/blob/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence/po"
)

func BlobToPo(blob *entity.Blob) *po.Blob {
	return &po.Blob{
		BlobID:     blob.BlobID,
		UserID:     blob.UserID,
		ObjectName: blob.ObjectName,
		BlobType:   blob.BlobType,
	}
}

func BlobToEntity(blob *po.Blob) *entity.Blob {
	return &entity.Blob{
		BlobID:     blob.BlobID,
		UserID:     blob.UserID,
		ObjectName: blob.ObjectName,
		BlobType:   blob.BlobType,
	}
}
