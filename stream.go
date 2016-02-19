package pr0gramm

import "errors"

type Stream struct {
  req ItemsRequest
  eof bool
}

func NewStream(req ItemsRequest) Stream {
  return Stream{req, false}
}

func (s *Stream) State() ItemsRequest {
  return s.req
}

func (s *Stream) Next() (*Items, error) {
  if ! s.More() {
    return nil, errors.New("Already at end of feed")
  }

  items, err := GetItems(s.req)
  if err == nil {
    s.eof = items.AtEnd

    if len(items.Items) > 0 {
      s.req.Older = items.Items[len(items.Items) - 1].Id
    }
  }

  return &items, err
}

func (s Stream) More() bool {
  return !s.eof
}
