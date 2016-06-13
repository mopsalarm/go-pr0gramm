package pr0gramm

type ContentType int
type ContentTypes []ContentType

const (
	SFW  ContentType = 1
	NSFW ContentType = 2
	NSFL ContentType = 4
	AUDIO ContentType = 8
)

var AllContentTypes = ContentTypes{SFW, NSFW, NSFL, AUDIO}

func (types ContentTypes) AsFlags() int {
	var result int
	for _, val := range types {
		result = result | int(val)
	}

	if result == 0 {
		result = 1
	}

	return result
}

func ToContentTypes(flags int) ContentTypes {
	var result ContentTypes

	if flags&int(SFW) != 0 {
		result = append(result, SFW)
	}
	if flags&int(NSFW) != 0 {
		result = append(result, NSFW)
	}
	if flags&int(NSFL) != 0 {
		result = append(result, NSFL)
	}
	if flags&int(AUDIO) != 0 {
		result = append(result, AUDIO)
	}

	return result
}
