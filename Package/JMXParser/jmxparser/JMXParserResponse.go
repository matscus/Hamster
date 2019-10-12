package jmxparser

type JMXParserResponse struct {
	TreadGroupName   string
	TreadGroupParams []TreadGroupParams
}
type TreadGroupParams struct {
	Name   string
	Values string
}
