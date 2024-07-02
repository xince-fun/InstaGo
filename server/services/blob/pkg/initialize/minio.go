package initialize

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/notification"
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

	for _, bucket := range c.Buckets {
		exists, err := mc.BucketExists(context.Background(), bucket)
		if err != nil {
			klog.Fatalf("check bucket exists failed %v", err)
		}
		if !exists {
			err = mc.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{Region: "cn-north-1"})
			if err != nil {
				klog.Fatalf("make bucket failed %v", err)
			}
		}
		err = mc.SetBucketPolicy(context.Background(), bucket, fmt.Sprintf(policy, bucket))
		if err != nil {
			klog.Fatalf("set bucket policy failed %v", err)
		}

		config := notification.Configuration{}

		topicArn := notification.NewArn("minio", "sqs", "", "instago-blob", "amqp")
		topicConfig := notification.NewConfig(topicArn)
		topicConfig.AddEvents(notification.ObjectCreatedAll, notification.ObjectRemovedAll)

		config.AddQueue(topicConfig)

		if err = mc.SetBucketNotification(context.Background(), bucket, config); err != nil {
			klog.Fatalf("set bucket notification failed %v", err)
		}
	}

	return mc
}
