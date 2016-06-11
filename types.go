package pr0gramm

import "time"

type Id uint64

type Response struct {
	Timestamp    Timestamp     `json:"ts"`
	ResponseTime time.Duration `json:"rt"`
	QueryCount   uint          `json:"qt"`
}

type Item struct {
	Id        Id        `json:"id"`
	Promoted  Id        `json:"promoted"`
	Up        uint      `json:"up"`
	Down      uint      `json:"down"`
	Created   Timestamp `json:"created"`
	Image     string    `json:"image"`
	Thumbnail string    `json:"thumb"`
	Fullsize  string    `json:"fullsize"`
	Source    string    `json:"source"`
	Flags     int       `json:"flags"`
	User      string    `json:"user"`
	Mark      int       `json:"mark"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Audio     bool      `json:"audio"`
}

type Items struct {
	Response

	Error   *string `json:"error"`
	AtEnd   bool    `json:"atEnd"`
	AtStart bool    `json:"atStart"`
	Items   []Item  `json:"items"`
}

type Tag struct {
	Id         Id      `json:"id"`
	Confidence float64 `json:"confidence"`
	Tag        string  `json:"tag"`
}

type BaseComment struct {
	Id      Id        `json:"id"`
	Created Timestamp `json:"created"`
	Up      uint      `json:"up"`
	Down    uint      `json:"down"`
	Content string    `json:"content"`
}

type Comment struct {
	BaseComment

	Mark    int    `json:"mark"`
	Parent  uint64 `json:"parent"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type UserComment struct {
	// normally i would use BaseComment, but the api sends a lot of stuff as strings.
	// BaseComment

	Id      Id        `json:"id,string"`
	Created Timestamp `json:"id"`
	Up      uint      `json:"up,string"`
	Down    uint      `json:"down,string"`
	Content string    `json:"content"`

	Thumbnail string `json:"thumb"`
	ItemId    uint64 `json:"itemId,string"`
}

type ItemThumb struct {
	Id        Id     `json:"id,string"`
	Thumbnail string `json:"thumb"`
}

type ItemInfo struct {
	Response

	Comments []Comment `json:"comments"`
	Tags     []Tag     `json:"tags"`
}

type InnerUserInfo struct {
	Id         Id        `json:"id"`
	Mark       int       `json:"mark"`
	Name       string    `json:"name"`
	Registered Timestamp `json:"registered"`
	Score      int       `json:"score"`
	Admin      uint      `json:"admin"`
	Banned     uint      `json:"banned"`
}

type UserInfo struct {
	Response

	CommentCount uint          `json:"commentCount"`
	Comments     []UserComment `json:"comments"`

	FollowCount uint `json:"followCount"`
	Following   bool `json:"following"`

	LikeCount      uint        `json:"likeCount"`
	LikesArePublic bool        `json:"likesArePublic"`
	Likes          []ItemThumb `json:"likes"`

	TagCount uint `json:"tagCount"`

	UploadCount uint        `json:"uploadCount"`
	Uploads     []ItemThumb `json:"uploads"`

	User InnerUserInfo `json:"user"`
}
