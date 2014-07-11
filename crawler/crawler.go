package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var startingpoint = "http://api.rottentomatoes.com/api/public/v1.0.json"
var apikey string
var count = 10

type LinkHolder struct {
	Links map[string]string `json:"links"`
}

type LinkListener struct {
	index          int
	linkChannel    chan string
	successChannel chan bool
}

func main() {
	apikey = os.Getenv("APIKEY")

	if apikey == "" {
		log.Fatal("Must provide a valid apikey")
	}
	if url := os.Getenv("URL"); url != "" {
		startingpoint = url
	}
	if cnt := os.Getenv("SIZE"); cnt != "" {
		tmp, err := strconv.Atoi(cnt)
		if err == nil {
			count = tmp
		}
	}

	runAndWait()
}

func runAndWait() {
	var wg sync.WaitGroup
	wg.Add(count)

	linkChannel := make(chan string, 100)
	successChannel := make(chan bool, count)
	startListeners(&wg, count, linkChannel, successChannel)
	startCounter(successChannel)
	linkChannel <- extractUrl(startingpoint)

	wg.Wait()
	fmt.Println("Done")
}

func startCounter(successChannel chan bool) {
	go func() {
		counter := 0
		go func() {
			<-time.After(5 * time.Second)
			log.Println("Final Count:", counter)
			os.Exit(0)
		}()
		for {
			timeout := getTimeoutChannelInSeconds()
			select {
			case <-successChannel:
				counter++
				log.Println("Count:", counter)
			case <-timeout:
				return
			}
		}
	}()
}

func startListeners(wg *sync.WaitGroup, count int, linkChannel chan string, successChannel chan bool) {
	for i := 0; i < count; i++ {
		curr := i
		go func() {
			log.Println("Starting Listener:", curr)
			listener := &LinkListener{index: curr, linkChannel: linkChannel, successChannel: successChannel}
			listener.StartListening()
			log.Println("Finished Listener:", curr)
			wg.Done()
		}()
	}
}

func (listener *LinkListener) StartListening() {
	for {
		timeout := getTimeoutChannelInSeconds()
		select {
		case url := <-listener.linkChannel:
			//log.Println("  ", listener.index, "retrieving url:", url)
			log.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				log.Println("  ", listener.index, "Caught Error Getting URL:", err)
			} else {
				defer resp.Body.Close()
				listener.parseBody(resp.Body)
			}
		case <-timeout:
			return
		}
	}
}

func (listener *LinkListener) parseBody1(body io.Reader) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("  ", listener.index, "Caught Error Parsing Body:", err)
	}
	var l LinkHolder

	json.Unmarshal(b, &l)
	listener.processLinks1(l.Links)
}
func (listener *LinkListener) parseBody(body io.Reader) {
	dec := json.NewDecoder(body)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			//log.Println("  ", listener.index, "Caught Error Parsing Body:", err)
			return
		}
		listener.navigateJsonObject(v)
	}
}
func (listener *LinkListener) navigateJsonObject(v map[string]interface{}) {
	for k := range v {
		//log.Println("  ", listener.index, "found key:", k)

		switch t := v[k].(type) {
		case map[string]interface{}:
			if k == "links" {
				listener.processLinks(t)
			} else {
				listener.navigateJsonObject(t)
			}
		case []interface{}:
			listener.navigateJsonArray(t)
		default: // ignore
		}
	}
}
func (listener *LinkListener) navigateJsonArray(arr []interface{}) {
	for _, v := range arr {
		//log.Println("  ", listener.index, "found array value:", v)

		switch t := v.(type) {
		case map[string]interface{}:
			listener.navigateJsonObject(t)
		case []interface{}:
			listener.navigateJsonArray(t)
		default: // ignore
		}
	}
}

func (listener *LinkListener) processLinks1(links map[string]string) {
	//log.Println("  ", listener.index, "  found", len(links), "more links")
	for _, v := range links {
		//log.Println("  ", listener.index, "  link:", v)
		listener.linkChannel <- extractUrl(v)
		listener.successChannel <- true
	}
}

func (listener *LinkListener) processLinks(links map[string]interface{}) {
	//log.Println("  ", listener.index, "  found", len(links), "more links")
	for _, v := range links {
		//log.Println("  ", listener.index, "  link:", v)
		url := extractUrl(v.(string))
		go func() {
			listener.linkChannel <- url
			listener.successChannel <- true
		}()
	}
}

func extractUrl(link string) string {
	if idx := strings.Index(link, "?"); idx == -1 {
		return fmt.Sprintf("%v?apikey=%v&", link, apikey)
	} else {
		return fmt.Sprintf("%v&apikey=%v&", link, apikey)
	}
}
func getTimeoutChannelInSeconds() chan bool {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(5 * time.Second)
		timeout <- true
	}()
	return timeout
}

// 0) start up a bounded number of goroutines to listen on link channel

// 0.1) put first url on the channel

// unique filter sit in between

// worker

// 		1) Go Get URL

// 		2) Iterate through links, passing each one to a queue/channel
