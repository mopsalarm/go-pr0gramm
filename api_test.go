package pr0gramm

import (
  "testing"
)

func TestGetItems(t *testing.T) {
  req := NewItemsRequest().WithUser("Mopsalarm")
  response, err := GetItems(req)
  if err != nil {
    t.Error(err)
    return
  }

  t.Log(response)
}

func TestGetItem(t *testing.T) {
  var id uint64 = 1143703
  response, err := GetItemInfo(id)
  if err != nil {
    t.Error(err)
    return
  }

  t.Log(response)
}

func TestGetUser(t *testing.T) {
  response, err := GetUserInfoSfw("Bolok")
  if err != nil {
    t.Error(err)
    return
  }

  t.Log(response)
}

