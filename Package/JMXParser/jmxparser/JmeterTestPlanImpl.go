package jmxparser

import (
	"strings"
)

var (
	paramsName = map[string]string{"ThreadGroup.num_threads": "threads",
		"ThreadGroup.ramp_time": "rampup", "ThreadGroup.duration": "duration",
		"ThreadGroup.delay": "delay"}
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
		treadGroupName := jmx.HashTree.HashTree.ThreadGroup[i].Testname
		params := make([]TreadGroupParams, 0, largs)
		for _, vl := range jmx.HashTree.HashTree.ThreadGroup[i].StringProp {
			for k, v := range resParams {
				if strings.Contains(vl.Text, k) {
					paramTypeName, ok := paramsName[vl.Name]
					if ok {
						params = append(params, TreadGroupParams{ParamType: paramTypeName, Name: k, Values: v})
					} else {
						params = append(params, TreadGroupParams{ParamType: vl.Name, Name: k, Values: v})
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
							if strings.Contains(throughputName, k) {
								params = append(params, TreadGroupParams{ParamType: "TPS", Name: k, Values: v})
							}
							// lUniform := len(jmx.HashTree.HashTree.HashTree[i].HashTree[i1].UniformRandomTimer.StringProp)
							// for i2 := 0; i2 < lUniform; i2++ {
							// 	if jmx.HashTree.HashTree.HashTree[i].HashTree[i1].UniformRandomTimer.StringProp[i2].Text != "" {
							// 		uniformName := jmx.HashTree.HashTree.HashTree[i].HashTree[i1].UniformRandomTimer.StringProp[i2].Text
							// 		if strings.Contains(uniformName, k) {
							// 			params = append(params, TreadGroupParams{ParamType: "UniformRandomTimer", Name: k, Values: v})
							// 		}
							// 	}
							// }
						}
						break
					}
				}
				break
			}
		}
		if len(params) > 0 {
			res = append(res, JMXParserResponse{TreadGroupName: treadGroupName, TreadGroupParams: params})
		}
	}
	for i := 0; i < ltg; i++ {
		treadGroupName := jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup[i].Testname
		params := make([]TreadGroupParams, 0, largs)
		for _, vl := range jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup[i].StringProp {
			for k, v := range resParams {
				if strings.Contains(vl.Text, k) {
					params = append(params, TreadGroupParams{ParamType: vl.Name, Name: k, Values: v})
				}
			}
		}
		if len(params) > 0 {
			res = append(res, JMXParserResponse{TreadGroupName: treadGroupName, TreadGroupParams: params})
		}
	}
	return res, nil
}
