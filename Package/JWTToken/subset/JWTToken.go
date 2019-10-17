package subset

type Token interface {
	Check() bool
	Generate(user string, projects []string) (string, error)
}
