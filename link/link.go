package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type LinkCrawlResponse struct {
	Links map[string]string `json:"links"`
}

type Listener struct {
	id                  int
	linkChannel         chan string
	completedUrlChannel chan string
}

var apikey string = "wqm4rcmdm5b3pk4jk9497hju"
var startURL string = "http://api.rottentomatoes.com/api/public/v1.0/lists.json?apikey=wqm4rcmdm5b3pk4jk9497hju"
var threadCount int = 10

func main() {
	// get os vars

	// run and wait
	runAndWait()
}

func runAndWait() {
	var wg sync.WaitGroup
	wg.Add(threadCount)

	linkChannel := make(chan string, threadCount*20)
	completedUrlChannel := make(chan string, threadCount)
	startListeners(&wg, threadCount, linkChannel, completedUrlChannel)
	startCompletedUrlChannelHandler(completedUrlChannel)
	linkChannel <- startURL

	wg.Wait()
	log.Println("All Go Routines Finished")
}

// response channel
// success/count channel

func startListeners(wg *sync.WaitGroup, count int, linkChannel chan string, completedUrlChannel chan string) {
	for i := 0; i < count; i++ {
		currentGoRoutine := i
		go func() {
			listener := &Listener{id: currentGoRoutine, linkChannel: linkChannel, completedUrlChannel: completedUrlChannel}
			log.Println("[", listener.id, "] starting listener")
			listener.listen()
			log.Println("[", listener.id, "] finished listener")
			wg.Done()
		}()
	}
}

func startCompletedUrlChannelHandler(completedUrlChannel chan string) {
	go func() {
		counter := 0
		for {
			select {
			case url := <-completedUrlChannel:
				counter++
				log.Println("[", counter, "] Completed Another URL:", url)
			case <-time.After(5 * time.Second):
				log.Println("completion channel timeout")
				return
			}
		}
	}()
}

func (l *Listener) listen() {
	for {
		select {
		case url := <-l.linkChannel:
			l.processUrl(url)
		case <-time.After(5 * time.Second):
			log.Println("[", l.id, "] timeout")
			return
		}
	}
}

func (l *Listener) processUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[", l.id, "] ERROR getting URL: [", url, "],", err)
		return
	}
	defer resp.Body.Close()
	l.completedUrlChannel <- url
	l.parseBody(resp.Body)
}

func (l *Listener) parseBody(body io.Reader) {
	dec := json.NewDecoder(body)

	for {
		var linkresponse LinkCrawlResponse
		if err := dec.Decode(&linkresponse); err != nil {
			if err != io.EOF {
				log.Println("[", l.id, "] could not decode link response:", err)
			}
			return
		} else {
			log.Println("[", l.id, "] decoded link response:", linkresponse)
			l.findLinks(&linkresponse)
		}
	}
}

func (l *Listener) findLinks(response *LinkCrawlResponse) {
	for k, v := range response.Links {
		if k != "alternate" {
			l.passOnNewLink(extractGoodUrl(v))
		}
	}
}

func (l *Listener) passOnNewLink(link string) {
	go func() {
		l.linkChannel <- link
	}()
}

func extractGoodUrl(link string) string {
	if idx := strings.Index(link, "?"); idx == -1 {
		return fmt.Sprintf("%v?apikey=%v&", link, apikey)
	} else {
		return fmt.Sprintf("%v&apikey=%v&", link, apikey)
	}
}
