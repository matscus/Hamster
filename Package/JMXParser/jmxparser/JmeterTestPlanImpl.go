package jmxparser

import (
	"strings"
)

var (
	paramsName = map[string]string{"ThreadGroup.num_threads": "Threads",
		"ThreadGroup.ramp_time": "RampUp", "ThreadGroup.duration": "Duration",
		"ThreadGroup.delay": "Delay", "TargetLevel": "Threads", "Hold": "Duration"}
)

func (jmx JmeterTestPlan) GetTreadGroupsParams() ([]JMXParserResponse, error) {
	ltg := len(jmx.HashTree.HashTree.ThreadGroup)
	lbztg := len(jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup)
	cArgs := len(jmx.HashTree.HashTree.Arguments)
	resParams := make(map[string]string)
	for i := 0; i < cArgs; i++ {
		for _, k := range jmx.HashTree.HashTree.Arguments[i].CollectionProp.ElementProp {
			resParams[k.StringProp[0].Text] = k.StringProp[1].Text
		}
	}
	largs := len(resParams)
	res := make([]JMXParserResponse, 0, ltg+lbztg)
	for i := 0; i < ltg; i++ {
		threadGroupName := jmx.HashTree.HashTree.ThreadGroup[i].Testname
		params := make([]ThreadGroupParams, 0, largs)
		for _, vl := range jmx.HashTree.HashTree.ThreadGroup[i].StringProp {
			for k, v := range resParams {
				text := strings.Trim(vl.Text, "${}")
				if text == k {
					paramTypeName, ok := paramsName[vl.Name]
					if ok {
						params = append(params, ThreadGroupParams{Type: paramTypeName, Name: k, Value: v})
					} else {
						params = append(params, ThreadGroupParams{Type: vl.Name, Name: k, Value: v})
					}
				}
			}
		}
		for i := 0; i < len(jmx.HashTree.HashTree.HashTree); i++ {
			if jmx.HashTree.HashTree.HashTree[i].TestAction.Testname != "" {
				l := len(jmx.HashTree.HashTree.HashTree[i].HashTree)
				for i1 := 0; i1 < l; i1++ {
					if jmx.HashTree.HashTree.HashTree[i].HashTree[i1].ConstantThroughputTimer.Testname != "" {
						throughputName := jmx.HashTree.HashTree.HashTree[i].HashTree[i1].ConstantThroughputTimer.StringProp.Text
						for k, v := range resParams {
							text := strings.Trim(throughputName, "${}")
							if strings.Contains(text, k) {
								params = append(params, ThreadGroupParams{Type: "TPS", Name: k, Value: v})
							}
						}
						break
					}
				}
				break
			}
		}
		if len(params) > 0 {
			res = append(res, JMXParserResponse{ThreadGroupName: threadGroupName, ThreadGroupType: "DefaultThreadGroup", ThreadGroupParams: params})
		}
	}
	for i := 0; i < ltg; i++ {
		threadGroupName := jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup[i].Testname
		params := make([]ThreadGroupParams, 0, largs)
		for _, vl := range jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup[i].StringProp {
			for k, v := range resParams {
				text := strings.Trim(vl.Text, "${}")
				if text == k {
					paramTypeName, ok := paramsName[vl.Name]
					if ok {
						params = append(params, ThreadGroupParams{Type: paramTypeName, Name: k, Value: v})
					} else {
						params = append(params, ThreadGroupParams{Type: vl.Name, Name: k, Value: v})
					}
				}
			}
		}
		if len(params) > 0 {
			res = append(res, JMXParserResponse{ThreadGroupName: threadGroupName, ThreadGroupType: "BlazemeterConcurrencyThreadGroup", ThreadGroupParams: params})
		}
	}
	return res, nil
}
