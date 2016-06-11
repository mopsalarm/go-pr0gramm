package pr0gramm

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func apiGet(url string, target interface{}) error {
	response, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		io.Copy(ioutil.Discard, response.Body)
		response.Body.Close()
	}()

	return json.NewDecoder(response.Body).Decode(target)
}
