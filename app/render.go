package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

var md = goldmark.New(goldmark.WithExtensions(meta.Meta))

type PostDate time.Time

func (t PostDate) String() string {
	return time.Time(t).Format("02 Jan 2006")
}

type Post struct {
	Name            string
	Title           string
	Url             string
	PublicationDate PostDate
	Content         template.HTML
	Tags            []string
}

func loadPost(path string) (*Post, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ctx := parser.NewContext()
	var out strings.Builder
	md.Convert(b, &out, parser.WithContext(ctx))

	m := meta.Get(ctx)

	title, _ := m["title"].(string)
	dateStr, _ := m["publication_date"].(string)
	date, _ := time.Parse("2006-01-02", dateStr)

	// read tags
	tagsInt, ok := m["tags"].([]interface{})
	tags := []string{}
	if ok {
		for _, tagInt := range tagsInt {
			tag, ok := tagInt.(string)
			if ok {
				tags = append(tags, tag)
			}
		}
	}

	name := strings.TrimSuffix(filepath.Base(path), ".md")

	return &Post{
		Title:           title,
		Url:             "/post/" + name + ".html",
		PublicationDate: PostDate(date),
		Content:         template.HTML(out.String()),
		Tags:            tags,
		Name:            name,
	}, nil
}

func loadPostsInDir(contentDir string) ([]*Post, error) {
	var posts []*Post
	files, _ := filepath.Glob(filepath.Join(contentDir, "*.md"))
	for _, path := range files {
		post, err := loadPost(path)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}
