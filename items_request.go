package pr0gramm

const (
  SFW ContentType = 1
  NSFW ContentType = 2
  NSFL ContentType = 4
)

type ItemsRequest struct {
  Older, Newer, Around int
  Tags, User, Likes    string
  Flags                ContentTypes
}

func NewItemsRequest() ItemsRequest {
  return ItemsRequest{Flags: ContentTypes{SFW}}
}

func (req ItemsRequest) WithOlderThan(id int) ItemsRequest {
  req.Older = id
  req.Newer = 0
  req.Around = 0
  return req
}

func (req ItemsRequest) WithNewerThan(id int) ItemsRequest {
  req.Older = id
  req.Newer = 0
  req.Around = 0
  return req
}

func (req ItemsRequest) WithAround(id int) ItemsRequest {
  req.Older = 0
  req.Newer = 0
  req.Around = id
  return req
}

func (req ItemsRequest) WithTag(tag string) ItemsRequest {
  req.Tags = tag
  return req
}

func (req ItemsRequest) WithUser(user string) ItemsRequest {
  req.User = user
  return req
}

func (req ItemsRequest) WithLikes(user string) ItemsRequest {
  req.Likes = user
  return req
}

func (req ItemsRequest) WithFlags(flags []ContentType) ItemsRequest {
  req.Flags = flags
  return req
}

func (types ContentTypes) AsFlags() int {
  var result int
  for _, val := range types {
    result = result | int(val)
  }

  if (result == 0) {
    result = 1
  }

  return result
}