package pr0gramm

import (
  "strings"
  "net/url"
  "strconv"
  "fmt"
)

type ContentType int
type ContentTypes []ContentType

func makeUrl(path string) string {
  return "http://pr0gramm.com/api/" + strings.TrimLeft(path, "/")
}

func GetItems(req ItemsRequest) (Items, error) {
  query := make(url.Values)
  query.Set("flags", strconv.Itoa(req.Flags.AsFlags()))

  if req.Older > 0 {
    query.Set("older", strconv.Itoa(req.Older))
  }

  if req.Newer > 0 {
    query.Set("newer", strconv.Itoa(req.Newer))
  }

  if req.Around > 0 {
    query.Set("id", strconv.Itoa(req.Around))
  }

  if req.Tags != "" {
    query.Set("tags", req.Tags)
  }

  if req.User != "" {
    query.Set("user", req.User)
  }

  if req.Likes != "" {
    query.Set("likes", req.Likes)
  }

  uri := makeUrl("/items/get?" + query.Encode())

  var response Items
  err := apiGet(uri, &response)
  return response, err
}

func GetItemInfo(id uint64) (ItemInfo, error) {
  uri := makeUrl(fmt.Sprintf("/items/info?itemId=%d", id))

  var response ItemInfo
  err := apiGet(uri, &response)
  return response, err
}
