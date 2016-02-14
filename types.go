package pr0gramm

type Response struct {
  Timestamp    Timestamp `json:"ts"`
  ResponseTime uint `json:"rt"`
  QueryCount   uint `json:"qt"`
}

type Item struct {
  Id        uint64
  Promoted  uint64
  Up        uint
  Down      uint
  Created   Timestamp
  Image     string
  Thumbnail string `json:"thumb"`
  Fullsize  string
  Source    string
  Flags     int
  User      string
  Mark      int
}

type Items struct {
  Response

  Error   string
  AtEnd   bool `json:"atEnd"`
  AtStart bool `json:"atStart"`
  Items   []Item
}

type Tag struct {
  Id         uint64
  Confidence float64
  Tag        string
}

type BaseComment struct {
  Id      uint64
  Created Timestamp
  Up      uint
  Down    uint
  Content string
}

type Comment struct {
  BaseComment

  Mark    int
  Parent  uint64
  Name    string
  Content string
}

type UserComment struct {
  // normally i would use BaseComment, but the api sends a lot of stuff as strings.
  // BaseComment

  Id        uint64 `json:"id,string"`
  Created   Timestamp `json:"id"`
  Up        uint `json:"up,string"`
  Down      uint `json:"down,string"`
  Content   string

  Thumbnail string `json:"thumb"`
  ItemId    uint64 `json:"itemId,string"`
}

type ItemThumb struct {
  Id        uint64 `json:"id,string"`
  Thumbnail string `json:"thumb"`
}

type ItemInfo struct {
  Response

  Comments []Comment
  Tags     []Tag
}

type InnerUserInfo struct {
  Id         uint64
  Mark       int
  Name       string
  Registered Timestamp
  Score      int
  Admin      uint
  Banned     uint
}

type UserInfo struct {
  Response

  CommentCount   uint
  Comments       []UserComment

  FollowCount    uint
  Following      bool

  LikeCount      uint
  LikesArePublic bool
  Likes          []ItemThumb

  TagCount       uint

  UploadCount    uint
  Uploads        []ItemThumb

  User           InnerUserInfo
}
