package net

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pksetdev/model"
	"github.com/schollz/progressbar/v3"
)

func Download(application *model.Application, appbasePath string) {

	// Get the data
	url := application.Url + application.Filename
	resp, err := http.Get(url)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}
	defer resp.Body.Close()

	pathToAppArchive := appbasePath + "/" + application.Filename
	out, err := os.Create(pathToAppArchive)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}
	defer out.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"         > downloading",
	)

	// Write to both the file and the progress bar
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}
	verifyDownload(pathToAppArchive, application)
}

func verifyDownload(filePath string, application *model.Application) {
	downloadedFile, err := os.Open(filePath)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}
	defer downloadedFile.Close()

	hasher, errv := verifier(application.Algorithm)
	if errv != nil {
		pklog.CreateLog(pklog.FatalError,
			fmt.Sprintf("Hash algorithm %s is undefined or has not been implemented", application.Algorithm))
	}

	_, error := io.Copy(hasher, downloadedFile)
	if error != nil {
		pklog.CreateLog(pklog.FatalError, error.Error())
	}

	checksum := hasher.Sum(nil)
	checksumString := fmt.Sprintf("%x", checksum)
	isEqual := isHashEqual(checksumString, application.Hash)
	if isEqual {
		pklog.CreateLog(pklog.Information, fmt.Sprintf("hash %s: OK", application.Algorithm))
		pklog.CreateLog(pklog.Information, "downloaded successfully")
	} else {
		pklog.CreateLog(pklog.Information, fmt.Sprintf("hash %s: FAILED", application.Algorithm))
		pklog.CreateLog(pklog.Information, "downloaded unsuccessfully")
		errr := os.Remove(filePath)
		if errr != nil {
			pklog.CreateLog(pklog.Error, errr.Error())
		}
		pklog.CreateLog(pklog.Error, fmt.Sprintf("filePath %s", filePath))
	}
}

func verifier(algorithm string) (hash hash.Hash, err error) {
	alg := strings.ToLower(algorithm)

	switch alg {
	case "md5":
		return md5.New(), nil
	case "sha256":
		return sha256.New(), nil
	case "sha512":
		return sha512.New(), nil
	default:
	}

	return nil, errors.New("hash algorithm is undefined or has not been implemented")
}

func isHashEqual(checksum, checksumConfig string) bool {

	return checksum == checksumConfig
}
