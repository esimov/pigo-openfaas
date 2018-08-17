package function

import (
	"fmt"
	"net/http"
	"strings"
	"io"
	"os"
	"io/ioutil"
	"log"
)

// Handle a serverless request
func Handle(req []byte) string {
	inputUrl := strings.TrimSpace(string(req))
	res, err := http.Get(inputUrl)
	if err != nil {
		return fmt.Sprintf("Unable to download image file from URI: %s, status %d", inputUrl, res.Status)
	}
	defer res.Body.Close()

	tmpfile, err := ioutil.TempFile("/tmp", "image")
	if err != nil {
		log.Fatal("Unable to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = io.Copy(tmpfile, res.Body)
	if err != nil {
		return fmt.Sprintf("Unable to copy the source URI to the destionation file")
	}
	return
}