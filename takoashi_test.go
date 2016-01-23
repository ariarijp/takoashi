package takoashi

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFilename(t *testing.T) {
	func() {
		fn := getFilename("http://example.com/foo/bar/baz")
		assert.Equal(t, "baz", fn)
	}()

	func() {
		fn := getFilename("http://example.com/foo/bar/baz.tgz")
		assert.Equal(t, "baz.tgz", fn)
	}()
}

func TestCreateFile(t *testing.T) {
	res, err := http.Get("http://example.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fn := "./test"
	createFile(fn, res)
	defer os.Remove(fn)

	assert.True(t, Exists(fn))
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
