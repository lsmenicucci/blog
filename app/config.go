package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ContentDir  string `json:"content_dir"`
	TemplateDir string `json:"template_dir"`
	StaticDir   string `json:"static_dir"`
}

func (c Config) TemplatePath(name string) string {
	return fmt.Sprintf("%s/%s.html", c.TemplateDir, name)
}

var config Config

func LoadConfig(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &config)
}
