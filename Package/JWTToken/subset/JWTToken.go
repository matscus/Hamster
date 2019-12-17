package subset

//Token -
type Token interface {
	Check() bool
	Generate(role string, user string, projects []string) (tokenstring string, err error)
	GenerateTemp(role string, user string, projects []string) (tokenstring string, err error)
}
