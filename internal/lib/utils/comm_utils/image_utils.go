package comm_utils

import (
	"errors"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/google/uuid"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetLocalImageSize(imgPath string) (int64, int64, error) {
	file, err := os.Open(imgPath)
	defer file.Close()
	if err != nil {
		return 0, 0, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}
	bounds := img.Bounds()
	return int64(bounds.Max.X), int64(bounds.Max.Y), nil
}

func GetRemoteImageSize(imgUrl string) (int64, int64, error) {
	parsedUrl, err := url.Parse(imgUrl)
	if err != nil {
		return 0, 0, err
	}

	ss := strings.Split(parsedUrl.Path, ".")
	originalExt := ss[len(ss)-1]

	tmpFile, err := downloadFile(imgUrl, originalExt)
	if err != nil {
		return 0, 0, err
	}

	file, err := os.Open(tmpFile)
	defer file.Close()
	if err != nil {
		return 0, 0, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}
	bounds := img.Bounds()

	defer func() {
		err = os.Remove(tmpFile)
		if err != nil {
			log.Error().Err(err).Msgf("failed to remove file %s", tmpFile)
		}
	}()

	return int64(bounds.Max.X), int64(bounds.Max.Y), nil
}

func downloadFile(url, originalExt string) (string, error) {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("received non 200 response code")
	}

	//Create an empty file
	fileName := fmt.Sprintf("%s.%s", uuid.New().String(), originalExt)
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
