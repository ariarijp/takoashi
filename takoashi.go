package takoashi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)

func getFilename(u string) string {
	_u, err := url.Parse(u)
	if err != nil {
		log.Fatal(fmt.Sprint("Failed parse as URL", u, "-", err))
	}

	return path.Base(_u.Path)
}

func createFile(fn string, res *http.Response) {
	out, err := os.Create(fn)
	if err != nil {
		log.Fatal(fmt.Sprint("Error while creating", fn, "-", err))
	}
	defer out.Close()

	n, err := io.Copy(out, res.Body)
	if err != nil {
		log.Fatal(fmt.Sprint("Error while downloading", fn, "-", err))
	}

	fmt.Println(res.Request.URL, res.Status, n, "bytes downloaded.")
}

func download(wait *sync.WaitGroup, u string) {
	fn := getFilename(u)

	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	createFile(fn, res)

	wait.Done()
}

func Run() {
	if len(os.Args) == 1 {
		log.Fatal("Missing required arguments")
	}

	urls := os.Args[1:]
	wait := new(sync.WaitGroup)

	for _, u := range urls {
		wait.Add(1)
		go download(wait, u)
	}

	wait.Wait()
}
