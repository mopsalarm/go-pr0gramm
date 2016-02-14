package pr0gramm

type Response struct {
  Error        string `json:"error"`
  Timestamp    Timestamp `json:"ts"`
  ResponseTime uint `json:"rt"`
  QueryCount   uint `json:"qt"`
}

type Item struct {
  Id        uint64 `json:"id"`
  Promoted  uint64 `json:"promoted"`
  Up        uint `json:"up"`
  Down      uint `json:"down"`
  Created   Timestamp `json:"created"`
  Image     string `json:"image"`
  Thumbnail string `json:"thumb"`
  Fullsize  string `json:"fullsize"`
  Source    string `json:"source"`
  Flags     int `json:"flags"`
  User      string `json:"user"`
  Mark      int `json:"thumb"`
}

type Items struct {
  Response

  AtEnd   bool `json:"atEnd"`
  AtStart bool `json:"atStart"`
  Items   []Item `json:"items"`
}

type Tag struct {
  Id         uint64
  Confidence float64
  Tag        string
}

type Comment struct {
  Id         uint64
  Confidence float64
  Created    Timestamp
  Up         uint
  Down       uint
  Mark       int
  Parent     uint64
  Name       string
  Content    string
}

type ItemInfo struct {
  Response

  Comments []Comment `json:"comments"`
  Tags     []Tag `json:"tags"`
}
