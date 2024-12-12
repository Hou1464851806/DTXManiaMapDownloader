package client

import (
	"DTXMapDownload/app/downloader"
	"DTXMapDownload/pkg/global"
	"DTXMapDownload/pkg/utils"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type Client struct {
	colly   *colly.Collector
	url     string
	songMap []string
}

func newBaseCollector(u string) *Client {
	return &Client{
		colly:   colly.NewCollector(),
		url:     u,
		songMap: make([]string, 0),
	}
}

func NewCollector(u string) *Client {
	c := newBaseCollector(u)
	c.colly.OnRequest(func(request *colly.Request) {
		fmt.Println("请求开始---------->")
	})
	c.colly.OnError(func(response *colly.Response, err error) {
		fmt.Println("发生错误：", err)
	})
	c.colly.OnResponse(func(response *colly.Response) {
		fmt.Println("请求完毕---------->")
		fmt.Println("解析开始---------->")
	})
	c.colly.OnScraped(func(response *colly.Response) {
		fmt.Println("解析完毕---------->")
	})
	return c
}

func (c *Client) Search(name string) {
	u := utils.SetQuery(global.Settings.SourceURL+"/search", "q", name)
	c.url = u
	c.getSongsInfo()
	c.Collect()
}

func (c *Client) getSongsInfo() {
	c.songMap = make([]string, 0)
	c.colly.OnHTML("#main", func(element *colly.HTMLElement) {
		element.ForEachWithBreak("div[class='post-body entry-content']", func(i int, e *colly.HTMLElement) bool {
			fmt.Printf("第%d首\n", i+1)
			text := utils.Beauty(e.Text)
			fmt.Println(text)
			downloadLink := ""
			e.ForEach("a", func(i int, e *colly.HTMLElement) {
				if i != 0 {
					downloadLink = e.Attr("href")
					fmt.Println(downloadLink)
				}
			})
			c.songMap = append(c.songMap, downloadLink)
			return true
		})
		fmt.Printf("共%d首\n", len(c.songMap))
	})
}

func (c *Client) Download(name string) {
	if global.Settings.GameSongsPath == "" {
		fmt.Println("Please set your game songs storage path first")
		return
	}
	log.Println(global.Settings.GameSongsPath)
	c.Search(name)
	if len(c.songMap) == 0 {
		fmt.Println("Get songs info failed, please try again")
		return
	}
	d := downloader.NewDownload(c.songMap[0])
	err := d.Download()
	if err != nil {
		log.Printf("download songs error: %v", err)
	}
}

func (c *Client) Collect() {
	err := c.colly.Visit(c.url)
	if err != nil {
		log.Printf("爬取网站%s时出现错误\n", err)
	}
}

func (c *Client) Config(key string, value string) {
	switch key {
	case "game":
		global.Settings.GameSongsPath = value
	}
	fmt.Printf("Game songs storage path is set to %s\n", global.Settings.GameSongsPath)
	err := global.Settings.Save()
	if err != nil {
		log.Printf("config save failed: %v", err)
	}
}
