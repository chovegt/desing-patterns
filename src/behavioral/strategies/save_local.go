package strategies

const (
	saveLocalText = "Save local: "
)

// strategies with Struct
type SaveLocal struct{}

func (s SaveLocal) Execute(text string) string {
	return saveLocalText + text
}

// strategies with funtion
func ExecuteSaveLocal(text string) string {
	return saveLocalText + text
}
