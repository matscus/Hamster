package subset

type Token interface {
	Generate(user string, projects []string) (string, error)
	Check() bool
}
