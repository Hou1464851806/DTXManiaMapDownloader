package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const SettingsName = "settings.json"

type Config struct {
	SourceURL     string `json:"source_url"`
	GameSongsPath string `json:"game_songs_path"`
}

func NewConfig() *Config {
	return &Config{}
}

func NewDefaultConfig() *Config {
	return &Config{
		SourceURL:     "http://approvedtx.blogspot.com",
		GameSongsPath: "",
	}
}

func (c *Config) Default() {
	c.SourceURL = "https://approvedtx.blogspot.com"
	c.GameSongsPath = ""
}

func (c *Config) Load() error {
	if _, err := os.Stat(SettingsName); os.IsNotExist(err) {
		c.Default()
		err = c.Save()
		if err != nil {
			return err
		}
		return nil
	}

	log.Printf("Loading config from %s\n", SettingsName)
	file, err := os.Open(SettingsName)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Save() error {
	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return err
	}
	err = os.WriteFile(SettingsName, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
