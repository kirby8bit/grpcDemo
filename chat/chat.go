package chat

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	sync "sync"

	"golang.org/x/net/context"
)

type Server struct {
	UnsafeChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: getThumbnail(in.Body)}, nil
}

func getThumbnail(someUrl string) string {
	if err := os.Mkdir("thumbnails", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	url1 := get_url(someUrl)
	urls := make([]string, 2)
	urls[0] = url1
	urls[1] = url1

	fmt.Println(urls)

	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go download_file(i, url, wg)
	}
	defer wg.Wait()
	return "check your thumbnails !"
}
func download_file(i int, url string, wg sync.WaitGroup) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	file, err := os.Create("thumbnails/picture" + strconv.Itoa(i) + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success!")
	wg.Done()
}
func get_url(urlVideo string) string {
	someVideo := urlVideo
	u, err := url.Parse(someVideo)
	if err != nil {
		log.Fatal(err)
	}
	videoId := someVideo[16:]
	u.Host = "img.youtube.com"
	return u.Scheme + "://" + u.Host + "/vi" + videoId + "/sddefault.jpg"
}
