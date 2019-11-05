package scn

import (
	"github.com/matscus/Hamster/Package/JMXParser/jmxparser"
)

type PreParseResponce struct {
	ThreadGroupName string   `json:"ThreadGroupName"`
	FailedParams    []string `json:"FailedParams"`
}
type ScriptCache struct {
	ScriptFile  []byte
	ParseParams []jmxparser.JMXParserResponse
}
