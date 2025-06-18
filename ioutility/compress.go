package ioutility

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
)

func ExtractZip(pathToZipFile string) error {

	directory, _ := splitPath(pathToZipFile)

	r, err := zip.OpenReader(pathToZipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	counter := 0

	for _, f := range r.File {
		fpath := filepath.Join(directory, f.Name)

		// Prevent ZipSlip vulnerability
		if !strings.HasPrefix(fpath, filepath.Clean(directory)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}

		if counter%50 == 0 {
			fmt.Print(".")
		}
		counter++
	}
	fmt.Println("")

	return nil
}

func ExtractTarGz(pathToTarGzFile string) error {

	directory, _ := splitPath(pathToTarGzFile)

	tarGzFile, err := readFile(pathToTarGzFile)
	if err != nil {
		pklog.CreateLog(pklog.Error, "cannot uncompressed the file: "+err.Error())
		return err
	}
	defer tarGzFile.Close()

	uncompressedStreamTarGz, err := gzip.NewReader(tarGzFile)
	if err != nil {
		pklog.CreateLog(pklog.Error, "tar.gz file cannot be extracted: "+err.Error())
		return err
	}
	defer uncompressedStreamTarGz.Close()

	if err1 := extractTarFile(uncompressedStreamTarGz, directory); err1 != nil {
		pklog.CreateLog(pklog.Error, "tar.gz file cannot be extracted: "+err1.Error())
		return err1
	}
	return nil
}

func extractTarFile(uncompressedStreamTarGz *gzip.Reader, destination string) error {
	tarFile := tar.NewReader(uncompressedStreamTarGz)

	counter := 0

	for {
		header, err := tarFile.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(destination, header.Name)

		if err1 := createExtractedTarStructure(header, target, tarFile); err1 != nil {
			return err1
		}

		if counter%50 == 0 {
			fmt.Print(". ")
		}
		counter++

	}
	fmt.Println("")

	return nil
}

func createExtractedTarStructure(header *tar.Header, target string, tarFile *tar.Reader) error {

	switch header.Typeflag {
	case tar.TypeDir:
		directoryPermission := os.FileMode(header.Mode)
		if err := os.MkdirAll(target, directoryPermission); err != nil {
			return err
		}
	case tar.TypeReg:
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		outFile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(header.Mode))
		if err != nil {
			return err
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, tarFile); err != nil {
			outFile.Close()
			return err
		}
	case tar.TypeSymlink:
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}

		if err := os.Symlink(header.Linkname, target); err != nil {
			return err
		}
	default:
		pklog.CreateLog(pklog.Warning, fmt.Sprintf("Skipping unsupported type: %v, name: %s", header.Typeflag, header.Name))

	}
	return nil
}
