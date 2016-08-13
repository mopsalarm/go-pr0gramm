package pr0gramm

type ItemsRequest struct {
	Older, Newer, Around Id
	Tags, User, Likes    string
	Random, Top          bool
	ContentTypes         ContentTypes
}

func NewItemsRequest() ItemsRequest {
	return ItemsRequest{ContentTypes: AllContentTypes}
}

func (req ItemsRequest) WithOlderThan(id Id) ItemsRequest {
	req.Older = id
	req.Newer = 0
	req.Around = 0
	return req
}

func (req ItemsRequest) WithNewerThan(id Id) ItemsRequest {
	req.Older = id
	req.Newer = 0
	req.Around = 0
	return req
}

func (req ItemsRequest) WithAround(id Id) ItemsRequest {
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

func (req ItemsRequest) WithRandom(random bool) ItemsRequest {
	req.Random = random
	return req
}

func (req ItemsRequest) WithTopOnly(topOnly bool) ItemsRequest {
	req.Top = topOnly
	return req
}

func (req ItemsRequest) WithFlags(flags []ContentType) ItemsRequest {
	req.ContentTypes = flags
	return req
}
