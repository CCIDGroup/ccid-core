package artifact

import (
	"fmt"
	"github.com/CCIDGroup/ccid-core/utils"
	"github.com/minio/minio-go/v6"
	"log"
)

type Storage struct {
	option *StorageOption
	client *minio.Client
}

type StorageOption struct {
	Endpoint string
	AccessKeyID string
	SecretAccessKey string
	UseSSl bool
	BucketName string
	Location string
}

func (s *Storage) InitStorage(so *StorageOption){
	s.option = so
	var err error
	s.client, err = minio.New(so.Endpoint, so.AccessKeyID, so.SecretAccessKey, so.UseSSl)
	if err != nil {
		log.Fatalln(err)
	}

	err = s.client.MakeBucket(so.BucketName, so.Location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := s.client.BucketExists(so.BucketName)
		if errBucketExists == nil && exists {
			utils.LogMsg(fmt.Sprintf("We already own %s\n",so.BucketName))
		} else {
			utils.LogError(err,"error when creating bucket")
		}
	} else {
		utils.LogMsg(fmt.Sprintf("Successfully created %s\n", so.BucketName))
	}
}

func (s *Storage) UploadArtifact(objectName,filePath string){
	// Upload the zip file with FPutObject
	n, err := s.client.FPutObject(s.option.BucketName, objectName, filePath, minio.PutObjectOptions{ContentType:"application/zip"})
	if err != nil {
		log.Fatalln(err)
	}
	utils.LogMsg(fmt.Sprintf("Successfully uploaded %s of size %d\n", objectName, n))
}


