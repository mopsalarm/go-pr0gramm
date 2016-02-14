package pr0gramm

import "testing"

func TestToContentTypes(t *testing.T) {
  if v := ToContentTypes(1); len(v) != 1 || v[0] != SFW {
    t.Errorf("ToContentTypes(1) returned %v", v)
  }

  if v := ToContentTypes(4); len(v) != 1 || v[0] != NSFL {
    t.Errorf("ToContentTypes(4) returned %v", v)
  }

  if v := ToContentTypes(3); len(v) != 2 || v[0] != SFW || v[1] != NSFW {
    t.Errorf("ToContentTypes(3) returned %v", v)
  }
}

func TestAsFlags(t *testing.T) {
  if flags := (ContentTypes{SFW}).AsFlags(); flags != 1 {
    t.Errorf("SFW.flags was %v, but expected 1", flags)
  }

  if flags := (ContentTypes{SFW, NSFW}).AsFlags(); flags != 3 {
    t.Errorf("SFW.flags was %v, but expected 3", flags)
  }

  if flags := (ContentTypes{SFW, NSFW, NSFL}).AsFlags(); flags != 7 {
    t.Errorf("SFW.flags was %v, but expected 7", flags)
  }
}
