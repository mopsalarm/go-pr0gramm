package pr0gramm

import (
  "net/http"
  "encoding/json"
)

func apiGet(url string, target interface{}) error {
  response, err := http.DefaultClient.Get(url)
  if err != nil {
    return err
  }

  defer response.Body.Close()
  return json.NewDecoder(response.Body).Decode(target)
}
