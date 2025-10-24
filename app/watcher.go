package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/net/websocket"
)

var watcherLogger = log.New(os.Stdout, "[ watcher ] ", 0)

type Watcher struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
	watcher *fsnotify.Watcher
}

func NewWatcher(paths []string) *Watcher {
	w, _ := fsnotify.NewWatcher()
	ws := &Watcher{clients: make(map[*websocket.Conn]bool), watcher: w}
	for _, p := range paths {
		watcherLogger.Printf("Watching: %s\n", p)	
		w.Add(p)
	}
	go ws.run()
	return ws
}

func (ws *Watcher) run() {
	for {
		select {
		case e := <-ws.watcher.Events:
			if e.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
				ws.mu.Lock()
				for c := range ws.clients {
					watcherLogger.Println("Sending reload event")
					websocket.Message.Send(c, "reload")
				}
				ws.mu.Unlock()
			}
		case err := <-ws.watcher.Errors:
			watcherLogger.Println(err)
		}
	}
}

func (ws *Watcher) HandleWS() http.Handler {
	return websocket.Handler(func(c *websocket.Conn) {
		ws.mu.Lock()
		ws.clients[c] = true
		ws.mu.Unlock()
		defer func() {
			ws.mu.Lock()
			delete(ws.clients, c)
			ws.mu.Unlock()
			c.Close()
		}()
		select {}
	})
}
