package s3

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/crypto-tester/terraform-backend/pkg/storage"
	"github.com/crypto-tester/terraform-backend/pkg/terraform"
)

const Name = "s3"

type S3Storage struct {
	client *minio.Client
	bucket string
}

func NewS3Storage(endpoint, bucket, accessKey, secretKey string, useSSL bool) (*S3Storage, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize minio client: %w", err)
	}

	if exists, err := client.BucketExists(context.Background(), bucket); err != nil {
		return nil, fmt.Errorf("failed to check for bucket: %w", err)
	} else if !exists {
		if err = client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("bucket does not exist and creation failed: %w", err)
		}
	}

	return &S3Storage{
		client: client,
		bucket: bucket,
	}, nil
}

func (s *S3Storage) GetName() string {
	return Name
}

func (s *S3Storage) SaveState(state *terraform.State) error {
	r := bytes.NewReader(state.Data)
	_, err := s.client.PutObject(context.Background(), s.bucket, getObjectName(state.ID), r, r.Size(), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	return err
}

func (s *S3Storage) GetState(id string) (*terraform.State, error) {
	state := &terraform.State{
		ID: id,
	}

	obj, err := s.client.GetObject(context.Background(), s.bucket, getObjectName(id), minio.GetObjectOptions{})
	if err != nil {
		return state, err
	}
	defer obj.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(obj); err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return nil, storage.ErrStateNotFound
		}
		return state, err
	}

	state.Data = buf.Bytes()
	return state, nil
}

func (s *S3Storage) DeleteState(id string) error {
	return s.client.RemoveObject(context.Background(), s.bucket, getObjectName(id), minio.RemoveObjectOptions{})
}

func getObjectName(id string) string {
	return fmt.Sprintf("%s.tfstate", id)
}
