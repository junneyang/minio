package main

import (
	"errors"
	"io"
	"os"
	"path"

	"github.com/minio-io/minio/pkgs/storage"
	es "github.com/minio-io/minio/pkgs/storage/encodedstorage"
)

func erasureGetList(config inputConfig) (io.Reader, error) {
	// do nothing
	return nil, errors.New("Not Implemented")
}

func erasureGet(config inputConfig, objectPath string) (io.Reader, error) {
	var objectStorage storage.ObjectStorage
	rootDir := path.Join(config.rootDir, config.storageDriver)
	objectStorage, err := es.NewStorage(rootDir, 10, 6, 1024*1024)
	if err != nil {
		return nil, err
	}
	object, err := objectStorage.Get(objectPath)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func erasurePut(config inputConfig, objectPath string, reader io.Reader) error {
	var err error
	rootDir := path.Join(config.rootDir, config.storageDriver)
	if err := os.MkdirAll(rootDir, 0700); err != nil {
		return err
	}
	var objectStorage storage.ObjectStorage
	if objectStorage, err = es.NewStorage(rootDir, 10, 6, 1024*1024); err != nil {
		return err
	}
	if err = objectStorage.Put(objectPath, reader); err != nil {
		return err
	}
	return nil
}
