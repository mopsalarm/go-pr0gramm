package pr0gramm

import (
  "fmt"
  "testing"
)

func TestGetItems(t *testing.T) {
  req := NewItemsRequest().WithUser("Mopsalarm")
  response, err := GetItems(req)
  if err != nil {
    t.Error(err)
    return
  }

  fmt.Println(response)
}

func TestGetItem(t *testing.T) {
  var id uint64 = 1143703
  response, err := GetItemInfo(id)
  if err != nil {
    t.Error(err)
    return
  }

  fmt.Println(response)
}

func TestGetUser(t *testing.T) {
  response, err := GetUserInfoSfw("Bolok")
  if err != nil {
    t.Error(err)
    return
  }

  fmt.Println(response)
}

