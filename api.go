package pr0gramm

import (
	"net/url"
	"strconv"
	"strings"
)

func (id Id) String() string {
	return strconv.Itoa(int(id))
}

func (sess *Session) GetItems(req ItemsRequest) (Items, error) {
	query := make(url.Values)
	query.Set("flags", strconv.Itoa(req.ContentTypes.AsFlags()))

	if req.Older > 0 {
		query.Set("older", req.Older.String())
	}

	if req.Newer > 0 {
		query.Set("newer", req.Newer.String())
	}

	if req.Around > 0 {
		query.Set("id", req.Around.String())
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

	var response Items
	err := sess.apiGET("/items/get", query, &response)
	return response, err
}

func (sess *Session) GetItemInfo(id Id) (ItemInfo, error) {
	query := make(url.Values)
	query.Set("itemId", strconv.Itoa(int(id)))

	var response ItemInfo
	err := sess.apiGET("/items/info", query, &response)
	return response, err
}

func (sess *Session) GetUserInfoSfw(user string) (UserInfo, error) {
	return sess.GetUserInfo(user, ContentTypes{SFW})
}

func (sess *Session) GetUserInfo(user string, flags ContentTypes) (UserInfo, error) {
	query := make(url.Values)
	query.Set("name", user)
	query.Set("flags", strconv.Itoa(flags.AsFlags()))

	var response UserInfo
	err := sess.apiGET("/profile/info", query, &response)
	return response, err
}

func (sess *Session) TagsAdd(itemId Id, tags []string) error {
	var body struct {
		ItemId Id     `form:"itemId"`
		Tags   string `form:"tags"`
	}

	body.ItemId = itemId
	body.Tags = strings.Join(tags, ",")

	return sess.apiPOST("/tags/add", nil, body, nil)
}
