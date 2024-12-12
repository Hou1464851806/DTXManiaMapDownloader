package downloader

import (
	"DTXMapDownload/pkg/global"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const actualTemplate = "https://drive.google.com/uc?id=%s&export=download"
const confirmTemplate = "https://drive.usercontent.google.com/download?id=%s&export=download&confirm=t&uuid=%s"

type DownloadInfo struct {
	raw      string
	actual   string
	confirm  string
	fileID   string
	fileUUID string
}

func NewDownload(raw string) *DownloadInfo {
	d := &DownloadInfo{
		raw: raw,
	}
	return d
}

func (d *DownloadInfo) extractFileID() error {
	startIndex := strings.Index(d.raw, "/d/") + len("/d/")
	endIndex := strings.Index(d.raw, "/view")
	if startIndex == -1 || endIndex == -1 {
		return errors.New("invalid google drive link")
	}
	d.fileID = d.raw[startIndex:endIndex]
	return nil
}

func (d *DownloadInfo) getActualLink() error {
	if d.fileID == "" {
		return errors.New("fileID does not exist")
	}
	d.actual = fmt.Sprintf(actualTemplate, d.fileID)
	return nil
}

func (d *DownloadInfo) Download() error {
	err := os.MkdirAll(global.RepoName, os.ModePerm)
	if err != nil {
		return err
	}
	err = d.extractFileID()
	if err != nil {
		return err
	}
	err = d.getActualLink()
	log.Println(d.actual)
	if err != nil {
		return err
	}
	resp, err := http.Get(d.actual)
	if err != nil {
		return err
	}
	// fmt.Println(resp)
	err = d.confirmDownloadLarge(resp)
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s.zip", uuid.New().String())
	filePath := filepath.Join(global.RepoName, fileName)
	err = DownloadFileWithProgress(d.confirm, filePath)
	if err != nil {
		return err
	}
	err = Unzip(global.Settings.GameSongsPath, filePath)
	if err != nil {
		return err
	}
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

func (d *DownloadInfo) confirmDownloadLarge(resp *http.Response) error {
	contentType := resp.Header.Get("Content-Type")
	log.Println(contentType)
	if strings.Index(contentType, "text/html") == -1 {
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	html := string(body)
	if keyIndex := strings.Index(html, "uuid"); keyIndex != -1 {
		html = html[keyIndex+len("uuid"):]
		re := regexp.MustCompile(`value="([a-fA-F0-9-]{36})"`)
		matches := re.FindStringSubmatch(html)
		if matches == nil {
			return errors.New("extract file uuid error")
		}
		d.fileUUID = matches[1]
	}
	d.confirm = fmt.Sprintf(confirmTemplate, d.fileID, d.fileUUID)
	log.Println(d.confirm)
	return nil
}
