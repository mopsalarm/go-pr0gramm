package pr0gramm

import "errors"

type Stream struct {
  req ItemsRequest
  eof bool
}

func NewStream(req ItemsRequest) Stream {
  return Stream{req, false}
}

func (s Stream) Next() (*Items, error) {
  if ! s.More() {
    return nil, errors.New("Already at end of feed")
  }

  items, err := GetItems(s.req)
  if err == nil {
    s.eof = items.AtEnd
  }

  return &items, err
}

func (s Stream) More() bool {
  return !s.eof
}
