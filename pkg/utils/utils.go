package utils

import (
	"log"
	"net/url"
	"strings"
)

func Beauty(raw string) string {
	raw = strings.Map(func(r rune) rune {
		switch r {
		case 160:
			return ' '
		case '\n':
			return ' '
		}
		return r
	}, raw)
	raw = strings.TrimFunc(raw, func(r rune) bool {
		return r == 10 || r == 32
	})
	return raw
}

func GetSongName(data string) string {
	endIndex := strings.Index(data, "/")
	if endIndex == -1 {
		return ""
	}
	name := strings.TrimSpace(data[:endIndex])
	return name
}

func SetQuery(raw string, key string, value string) string {
	q := url.Values{}
	q.Add(key, value)
	u, err := url.Parse(raw)
	if err != nil {
		log.Printf("parse raw url error: %v", err)
	}
	u.RawQuery = q.Encode()
	log.Println(u.String())
	return u.String()
}

func ContainsString(elem string, list []string) bool {
	for _, l := range list {
		if elem == l {
			return true
		}
	}
	return false
}

func CompleteToFullURL(raw string) string {
	if !strings.Contains(raw, "https://") && !strings.Contains(raw, "http://") {
		raw = "http://" + raw
	}
	return raw
}
