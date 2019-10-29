package scn

import (
	"mime/multipart"

	"github.com/matscus/Hamster/Package/JMXParser/jmxparser"
)

type PreParseResponce struct {
	ThreadGroupName string   `json:"ThreadGroupName"`
	FailedParams    []string `json:"FailedParams"`
}
type ScriptCache struct {
	ScriptFile  multipart.File
	ParseParams []jmxparser.JMXParserResponse
}
