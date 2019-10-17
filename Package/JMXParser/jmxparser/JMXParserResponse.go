package jmxparser

type JMXParserResponse struct {
	ThreadGroupName   string
	ThreadGroupType   string
	ThreadGroupParams []ThreadGroupParams
}
type ThreadGroupParams struct {
	Type  string
	Name  string
	Value string
}
