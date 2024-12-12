package downloader

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Unzip(dst string, src string) error {
	zipReader, err := zip.OpenReader(src)
	defer zipReader.Close()
	if err != nil {
		return err
	}
	if dst != "" {
		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return err
		}
	}
	for _, file := range zipReader.File {
		filePath := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(filePath, file.Mode())
			if err != nil {
				return err
			}
			continue
		}
		err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}
		err = copyFile(file, filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyFile(file *zip.File, path string) error {
	fr, err := file.Open()
	defer fr.Close()
	if err != nil {
		return err
	}
	fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
	defer fw.Close()
	if err != nil {
		return err
	}
	n, err := io.Copy(fw, fr)
	if err != nil {
		return err
	}
	log.Printf("Unzip success: %s, total: %d\n", path, n)
	return nil
}
