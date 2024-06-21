package initialize

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/xince-fun/InstaGo/server/services/blob/conf"
)

func InitMinio() *minio.Client {
	c := conf.GlobalServerConf.BucketConfig

	mc, err := minio.New(c.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKeyID, c.AccessSecret, ""),
		Secure: false,
	})
	if err != nil {
		klog.Fatalf("create minio client failed %v", err)
	}
	exists, err := mc.BucketExists(context.Background(), c.AvatarBucket)
	if err != nil {
		klog.Fatalf("check bucket exists failed %v", err)
	}
	if !exists {
		err = mc.MakeBucket(context.Background(), c.AvatarBucket, minio.MakeBucketOptions{Region: "cn-north-1"})
		if err != nil {
			klog.Fatalf("make bucket failed %v", err)
		}
	}
	policy := `{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "AWS": ["*"]
                },
                "Action": [
                    "s3:GetObject",
                    "s3:PutObject"
                ],
                "Resource": [
                    "arn:aws:s3:::%s/*"
                ]
            }
        ]
    }`
	err = mc.SetBucketPolicy(context.Background(), c.AvatarBucket, fmt.Sprintf(policy, c.AvatarBucket))
	if err != nil {
		klog.Fatalf("set bucket policy failed %v", err)
	}
	return mc
}
