package pr0gramm

import (
	"testing"
	"time"
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
	id := Id(1143703)
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

func TestStream(t *testing.T) {
	req := NewItemsRequest()

	Stream(req, ConsumeIf(func(item Item) bool {
		t.Logf("Checking: %f\n", time.Since(item.Created.Time).Minutes())
		return time.Since(item.Created.Time).Minutes() < 10
	}, func(item Item) error {
		t.Logf("Accepted: %#v\n", item)
		return nil
	}))
}
