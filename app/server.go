package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var logger = log.New(os.Stdout, "[ server ] ", 0)

func renderTemplate(w io.Writer, tmpl string, data map[string]any) {
	t := template.Must(template.ParseFiles(config.TemplatePath("base"), tmpl))
	data["dev"] = true
	t.ExecuteTemplate(w, "base.html", data)
}

var allPosts = map[string]*Post{}

func renderPostList(w io.Writer) error {
	posts, err := loadPostsInDir(config.ContentDir)
	if err != nil {
		return err
	}
	for _, post := range posts {
		allPosts[post.Name] = post
	}

	data := map[string]any{"Posts": posts}
	renderTemplate(w, config.TemplatePath("list"), data)
	return nil
}

func listHandler(w http.ResponseWriter, _ *http.Request) {
	logger.Println("Rendering post list")
	err := renderPostList(w)
	if err != nil {
		logger.Printf("Post list render error: %s\n", err)
	}
}

func renderPost(name string, w http.ResponseWriter) error {
	path := "content/" + name + ".md"

	post, err := loadPost(path)
	if err != nil {
		return err
	}
	allPosts[post.Name] = post

	data := map[string]any{
		"Post": post,
	}
	renderTemplate(w, config.TemplatePath("post"), data)

	return nil
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/post/")
	name = strings.TrimSuffix(name, ".html")

	logger.Printf("Rendering %s name\n", name)

	err := renderPost(name, w)
	if err != nil {
		logger.Printf("Error rendering '%s': %s\n", name, err)
		fmt.Fprintln(w, "Not found")
		return
	}
}

func serve() {
	w := NewWatcher([]string{config.ContentDir, config.StaticDir, config.TemplateDir})
	http.Handle("/ws", w.HandleWS())
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(config.StaticDir))))
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/post/", postHandler)

	hostname := "localhost"
	port := 8080

	addr := fmt.Sprintf("%s:%d", hostname, port)
	log.Printf("Listening at: http://%s\n", addr)
	http.ListenAndServe(addr, nil)
}
