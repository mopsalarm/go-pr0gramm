package pr0gramm

import (
	"net/http"
	"testing"
	"time"
)

func TestGetItems(t *testing.T) {
	req := NewItemsRequest().
		WithFlags(ContentTypes{SFW}).
		WithUser("Mopsalarm")

	response, err := NewSession(http.Client{Timeout: 10 * time.Second}).GetItems(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response)
}

func TestGetItem(t *testing.T) {
	id := Id(1143703)
	response, err := NewSession(http.Client{Timeout: 10 * time.Second}).GetItemInfo(id)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response)
}

func TestGetUser(t *testing.T) {
	response, err := NewSession(http.Client{Timeout: 10 * time.Second}).GetUserInfoSfw("Bolok")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(response)
}
