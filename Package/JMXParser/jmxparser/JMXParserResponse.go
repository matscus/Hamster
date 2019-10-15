package jmxparser

type JMXParserResponse struct {
	TreadGroupName   string
	TGType           string
	TreadGroupParams []TreadGroupParams
}
type TreadGroupParams struct {
	ParamType string
	Name      string
	Values    string
}
