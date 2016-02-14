package pr0gramm

import (
  "strings"
  "net/url"
  "strconv"
  "fmt"
)

func makeUrl(path string) string {
  return "http://pr0gramm.com/api/" + strings.TrimLeft(path, "/")
}

func (id Id) ToString() string {
  return strconv.FormatInt(int64(id), 10)
}

func ParseId(value string) Id {
  id, err := strconv.ParseInt(value, 10, 0)
  if err != nil {
    id = 0
  }

  return Id(id)
}

func GetItems(req ItemsRequest) (Items, error) {
  query := make(url.Values)
  query.Set("flags", strconv.Itoa(req.Flags.AsFlags()))

  if req.Older > 0 {
    query.Set("older", req.Older.ToString())
  }

  if req.Newer > 0 {
    query.Set("newer", req.Newer.ToString())
  }

  if req.Around > 0 {
    query.Set("id", req.Around.ToString())
  }

  if req.Tags != nil {
    query.Set("tags", *req.Tags)
  }

  if req.User != nil {
    query.Set("user", *req.User)
  }

  if req.Likes != nil {
    query.Set("likes", *req.Likes)
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

func GetUserInfoSfw(user string) (UserInfo, error) {
  return GetUserInfo(user, ContentTypes{SFW})
}

func GetUserInfo(user string, flags ContentTypes) (UserInfo, error) {
  query := make(url.Values)
  query.Set("name", user)
  query.Set("flags", strconv.Itoa(flags.AsFlags()))
  uri := makeUrl("/profile/info?" + query.Encode())

  var response UserInfo
  err := apiGet(uri, &response)
  return response, err
}
