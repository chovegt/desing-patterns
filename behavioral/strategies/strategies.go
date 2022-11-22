package strategies

import "errors"

var (
	mapFuntions = map[string]SaveFileFunc{
		"local": ExecuteSaveLocal,
		"cloud": ExecuteSaveCloud,
	}

	ErrFoo = errors.New("invalid type")
)

func Strategy(_type string) (SaveFileFunc, error) {
	f, ok := mapFuntions[_type]
	if !ok {
		return nil, ErrFoo
	}

	return f, nil
}

func New(text string, method SaveFile) WithStruct {
	return WithStruct{
		savefile: method,
		text:     text,
	}
}

type WithStruct struct {
	savefile SaveFile
	text     string
}

func (s WithStruct) Print() string {
	return s.savefile.Execute(s.text)
}
