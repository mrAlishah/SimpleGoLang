package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

// ImageStore is an interface to store laptop images
//The role of the store is to save the uploaded image file somewhere on the server or on the cloud
type ImageStore interface {
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error)
}

// DiskImageStore stores image on disk, and its info on memory
//will save image files to the disk, and store its information in memory
type DiskImageStore struct {
	//we need a mutex to handle concurrency
	mutex sync.RWMutex
	//we need the path of the folder to save laptop images
	imageFolder string
	//memory record infor store to map with the key is image ID and the value is some information of the image
	images map[string]*ImageInfo
}

// ImageInfo contains information of the laptop image
type ImageInfo struct {
	//the ID of the laptop
	LaptopID string
	//the type of the image (or its file extension: jpg/png)
	Type string
	//the path to the image file on disk
	Path string
}

// NewDiskImageStore returns a new DiskImageStore
//which is the image folder. And inside, we just need to initialize the images map:
func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}

// Save adds a new image to a laptop
func (store *DiskImageStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error) {
	imageID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	//call os.Create() to create the file. And we call imageData.WriteTo() to write the image data to the created file
	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageID, imageType)

	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	//If the file is written successfully, we need to save its information to the in-memory map. So we have to acquire the write lock of the store.
	store.mutex.Lock()
	defer store.mutex.Unlock()

	//Then we save the image information to the map with key is the ID of the image, and the value contains the laptop ID, the image type, and the path to the image file.
	store.images[imageID.String()] = &ImageInfo{
		LaptopID: laptopID,
		Type:     imageType,
		Path:     imagePath,
	}

	return imageID.String(), nil
}
