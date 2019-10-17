package jmxparser

//JMXParserResponse - struct from response parse
type JMXParserResponse struct {
	ThreadGroupName   string
	ThreadGroupType   string
	ThreadGroupParams []ThreadGroupParams
}

//ThreadGroupParams - struct for simple thread group params
type ThreadGroupParams struct {
	Type  string
	Name  string
	Value string
}
