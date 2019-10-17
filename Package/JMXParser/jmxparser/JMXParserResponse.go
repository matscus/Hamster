package jmxparser

type JMXParserResponse struct {
	ThreadGroupName   string
	ThreadGroupType   string
	ThreadGroupParams []ThreadGroupParams
}
type ThreadGroupParams struct {
	ParamType string
	Name      string
	Value     string
}
