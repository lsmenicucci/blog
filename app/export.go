package main

import (
	"io"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	os.MkdirAll(filepath.Dir(dst), 0755)
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

func copyDir(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(srcDir, path)
		return copyFile(path, filepath.Join(dstDir, rel))
	})
}

func Export(dist string) error {
	os.RemoveAll(dist)
	os.MkdirAll(dist, 0755)
	posts, err := loadPostsInDir(config.ContentDir)
	if err != nil {
		return err
	}
	for _, post := range posts {
		path := filepath.Join(dist, "post", post.Name+".html")
		os.MkdirAll(filepath.Dir(path), 0755)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		renderTemplate(f, config.TemplatePath("post"), map[string]any{"Post": post}, false)
		f.Close()
	}
	listFile := filepath.Join(dist, "index.html")
	f, err := os.Create(listFile)
	if err != nil {
		return err
	}
	renderTemplate(f, config.TemplatePath("list"), map[string]any{"Posts": posts}, false)
	f.Close()
	copyDir("static", filepath.Join(dist, "static"))
	return nil
}
