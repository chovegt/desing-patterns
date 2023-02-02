package strategies

// example with interface
type SaveFile interface {
	Execute(text string) string
}

// example with funtions
type SaveFileFunc func(text string) string
