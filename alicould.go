package main

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func AliCreateBlunker(EndPt string, AKey string, AKeySec string, BucketName string) bool {

	client, err := oss.New(EndPt, AKey, AKeySec)
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}
	exist, err := client.IsBucketExist(BucketName)

	if exist == false {
		err = client.CreateBucket(BucketName)
		if err != nil {
			fmt.Printf(err.Error())
			return false
		}

	}
	acl, err := client.GetBucketACL(BucketName)
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}

	fmt.Println(string(oss.ACLPublicRead))
	if acl.ACL != string(oss.ACLPublicRead) {
		client.SetBucketACL(BucketName, oss.ACLPublicRead)
	}
	return true
}

func AliPut(EndPt string, AKey string, AKeySec string, BucketName string, ObjKey string, Filename string) {
	client, err := oss.New(EndPt, AKey, AKeySec)
	if err != nil {
		fmt.Printf(err.Error())
	}
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	err = bucket.PutObjectFromFile(ObjKey, Filename)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}
