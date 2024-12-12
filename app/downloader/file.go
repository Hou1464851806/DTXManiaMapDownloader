package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ProgressDownloader struct {
	Writer      io.Writer
	TotalSize   int64
	Downloaded  int64
	LastUpdated time.Time
}

func (p *ProgressDownloader) Write(w []byte) (int, error) {
	n, err := p.Writer.Write(w)
	if err != nil {
		return n, err
	}
	p.Downloaded += int64(n)
	if time.Since(p.LastUpdated) > 100*time.Millisecond {
		p.updateProgress()
		p.LastUpdated = time.Now()
	}
	return n, nil
}

func (p *ProgressDownloader) updateProgress() {
	percentage := float64(p.Downloaded) / float64(p.TotalSize) * 100
	fmt.Printf("\rDownloading... %.2f%% (%d/%d bytes)", percentage, p.Downloaded, p.TotalSize)
}

func DownloadFileWithProgress(url string, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	totalSize := resp.ContentLength
	if totalSize <= 0 {
		return errors.New("unable to get file total size")
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	p := &ProgressDownloader{
		Writer:    file,
		TotalSize: totalSize,
	}
	fmt.Println("Start Download")
	_, err = io.Copy(p, resp.Body)
	if err != nil {
		return err
	}
	p.updateProgress()
	fmt.Println("\nDownload Complete")
	return nil
}
