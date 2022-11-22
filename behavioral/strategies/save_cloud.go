package strategies

const (
	saveCloudText = "Save cloud: "
)

// strategies with Struct
type SaveCloud struct{}

func (s SaveCloud) Execute(text string) string {
	return saveCloudText + text
}

// strategies with funtion
func ExecuteSaveCloud(text string) string {
	return saveCloudText + text
}
