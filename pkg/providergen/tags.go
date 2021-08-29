package providergen

import "github.com/fatih/structtag"

type FieldAttributes struct {
	IsKey      bool
	Searchable bool
	Lookup     bool
}

const tagKey = "hotcereal"

func ParseTag(s string) FieldAttributes {
	fa := FieldAttributes{}

	tgs, err := structtag.Parse(s)
	if err != nil {
		return fa
	}

	tag, err := tgs.Get(tagKey)
	if err != nil {
		return fa
	}

	for _, attr := range append(tag.Options, tag.Name) {
		switch attr {
		case "":
			fa.Lookup = true
		case "lookup":
			fa.Lookup = true
		case "searchable":
			fa.Searchable = true
		case "key":
			fa.IsKey = true
		}
	}

	return fa
}
