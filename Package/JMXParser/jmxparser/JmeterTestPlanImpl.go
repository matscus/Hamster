package jmxparser

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	paramsNames = map[string]string{"ThreadGroup.num_threads": "Threads",
		"ThreadGroup.ramp_time": "RampUp", "ThreadGroup.duration": "Duration",
		"ThreadGroup.delay": "Delay", "TargetLevel": "Threads", "Hold": "Duration"}
	tgOpenTags = []string{"<ThreadGroup", "<com.blazemeter.jmeter.threads.concurrency.ConcurrencyThreadGroup"}
)

//GetTreadGroupsParams - func to return slice jmeter thread groups params
func (jmx JmeterTestPlan) GetTreadGroupsParams(tempScripsBytes []byte) ([]JMXParserResponse, error) {
	ltg := len(jmx.HashTree.HashTree.ThreadGroup)
	lbztg := len(jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup)
	cArgs := len(jmx.HashTree.HashTree.Arguments)
	resParams := make(map[string]string)
	largs := len(resParams)
	res := make([]JMXParserResponse, 0, ltg+lbztg)
	for i := 0; i < cArgs; i++ {
		for _, k := range jmx.HashTree.HashTree.Arguments[i].CollectionProp.ElementProp {
			resParams[k.StringProp[0].Text] = k.StringProp[1].Text
		}
	}
	throughputIndexs := make([]ThroughputIndexs, 0, 0)
	//r := regexp.MustCompile(`name="throughput">\${__jexl2\(\${([^\"]+)}\)}</stringProp>`)
	r := regexp.MustCompile(`name="throughput">([^\"]+)</stringProp>`)
	thgRegexp := r.FindAllStringSubmatch(string(tempScripsBytes), -1)
	iterThg := 0
	for _, v := range thgRegexp {
		re := regexp.MustCompile(`\${_{0,2}P?\(?\w*\b,?([\w]+|\d+)\)?}`)
		tgRe := re.FindAllStringSubmatch(v[1], -1)
		throughputIndexs = append(throughputIndexs, ThroughputIndexs{Index: iterThg, Name: tgRe[0][1]})
		iterThg++
	}
	tgIndexs := make([]TreadGroupIndex, 0, 0)
	r = regexp.MustCompile(`ThreadGroup\" testname=\"(\w+)`)
	tgRegexp := r.FindAllStringSubmatch(string(tempScripsBytes), -1)
	iterTg := 0
	for _, v := range tgRegexp {
		if v[1] != "tearDown" && v[1] != "GlobalSetUp" {
			tgIndexs = append(tgIndexs, TreadGroupIndex{Index: iterTg, Name: v[1]})
			iterTg++
		}
	}
	thgTimer := make([]ConstantThroughputTimer, 0, ltg+lbztg)
	l := len(jmx.HashTree.HashTree.HashTree)
	for i := 0; i < l; i++ {
		if jmx.HashTree.HashTree.HashTree[i].TestAction.Testname != "" {
			ll := len(jmx.HashTree.HashTree.HashTree[i].HashTree)
			for ii := 0; ii < ll; ii++ {
				if jmx.HashTree.HashTree.HashTree[i].HashTree[ii].ConstantThroughputTimer.Testname != "" {
					thgTimer = append(thgTimer, jmx.HashTree.HashTree.HashTree[i].HashTree[ii].ConstantThroughputTimer)
				}
			}
		}
	}
	for i := 0; i < ltg; i++ {
		threadGroupName := jmx.HashTree.HashTree.ThreadGroup[i].Testname
		params := make([]ThreadGroupParams, 0, largs)
		for _, vl := range jmx.HashTree.HashTree.ThreadGroup[i].StringProp {
			paramValues := vl.Text
			if paramValues == "" && vl.Name != "ThreadGroup.delay" {
				paramTypeName, ok := paramsNames[vl.Name]
				if ok {
					params = append(params, ThreadGroupParams{Type: paramTypeName, Name: "", Value: ""})
				}
			} else {
				if strings.Contains(paramValues, "${__P(") {
					re := regexp.MustCompile(`\${__P\(\s*(.+?)\s*,\s*(\${.+?}|[0-9]+)\s*\)}`)
					paramsRegexp := re.FindAllStringSubmatch(paramValues, -1)
					for i := 0; i < len(paramsRegexp); i++ {
						if strings.Contains(paramsRegexp[i][2], "${") {
							text := strings.Trim(paramsRegexp[i][2], "${}")
							paramTypeName, ok := paramsNames[vl.Name]
							if ok {
								params = append(params, ThreadGroupParams{Type: paramTypeName, Name: paramsRegexp[i][1], Value: resParams[text]})
							} else {
								params = append(params, ThreadGroupParams{Type: vl.Name, Name: paramsRegexp[i][1], Value: resParams[text]})
							}
						} else {
							params = append(params, ThreadGroupParams{Type: vl.Name, Name: paramsRegexp[i][1], Value: paramsRegexp[i][2]})
						}
					}
				} else {
					_, err := strconv.Atoi(vl.Text)
					if err != nil {
						paramTypeName, ok := paramsNames[vl.Name]
						if ok && vl.Text != "" {
							params = append(params, ThreadGroupParams{Type: paramTypeName, Name: "", Value: ""})
						}
					} else {
						paramTypeName, ok := paramsNames[vl.Name]
						if ok && vl.Text != "" {
							params = append(params, ThreadGroupParams{Type: paramTypeName, Name: "", Value: vl.Text})
						}
					}
				}
			}
		}
		for i := 0; i < len(tgIndexs); i++ {
			if threadGroupName == tgIndexs[i].Name {
				text := throughputIndexs[i].Name
				for _, v := range thgTimer {
					r := regexp.MustCompile(`\${_{0,2}P?\(?\w*\b,?([\w]+|\d+)\)?}`)
					thgTimerText := r.FindAllStringSubmatch(v.StringProp.Text, -1)
					if len(thgTimerText) > 0 {
						if text == thgTimerText[0][1] {
							value := resParams[text]
							_, err := strconv.Atoi(text)
							if err != nil {
								params = append(params, ThreadGroupParams{Type: "TPS", Name: text, Value: value})
								break
							} else {
								params = append(params, ThreadGroupParams{Type: "TPS", Name: "", Value: text})
								break
							}
						}
					}
				}
			}
		}
		if len(params) > 0 {
			res = append(res, JMXParserResponse{ThreadGroupName: threadGroupName, ThreadGroupType: "DefaultThreadGroup", ThreadGroupParams: params})
		}
	}
	for i := 0; i < lbztg; i++ {
		threadGroupName := jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup[i].Testname
		params := make([]ThreadGroupParams, 0, largs)
		for _, vl := range jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup[i].StringProp {
			paramValues := vl.Text
			if paramValues == "" {
				paramTypeName, ok := paramsNames[vl.Name]
				if ok {
					params = append(params, ThreadGroupParams{Type: paramTypeName, Name: "", Value: ""})
				}
			} else {
				if strings.Contains(paramValues, "${__P(") {
					re := regexp.MustCompile(`\${__P\(\s*(.+?)\s*,\s*(\${.+?}|[0-9]+)\s*\)}`)
					paramsRegexp := re.FindAllStringSubmatch(paramValues, -1)
					for i := 0; i < len(paramsRegexp); i++ {
						if strings.Contains(paramsRegexp[i][2], "${") {
							text := strings.Trim(paramsRegexp[i][2], "${}")
							paramTypeName, ok := paramsNames[vl.Name]
							if ok {
								params = append(params, ThreadGroupParams{Type: paramTypeName, Name: paramsRegexp[i][1], Value: resParams[text]})
							} else {
								params = append(params, ThreadGroupParams{Type: vl.Name, Name: paramsRegexp[i][1], Value: resParams[text]})
							}
						} else {
							paramTypeName, ok := paramsNames[vl.Name]
							if ok {
								params = append(params, ThreadGroupParams{Type: paramTypeName, Name: paramsRegexp[i][1], Value: paramsRegexp[i][2]})
							} else {
								params = append(params, ThreadGroupParams{Type: vl.Name, Name: paramsRegexp[i][1], Value: paramsRegexp[i][2]})
							}
						}
					}
				} else if strings.Contains(paramValues, "${") {
					text := strings.Trim(paramValues, "${}")
					paramTypeName, ok := paramsNames[vl.Name]
					if ok {
						params = append(params, ThreadGroupParams{Type: paramTypeName, Name: text, Value: resParams[text]})
					} else {
						params = append(params, ThreadGroupParams{Type: vl.Name, Name: text, Value: resParams[text]})
					}
				} else {
					_, err := strconv.Atoi(vl.Text)
					if err != nil {
						paramTypeName, ok := paramsNames[vl.Name]
						if ok && vl.Text != "" {
							params = append(params, ThreadGroupParams{Type: paramTypeName, Name: "", Value: ""})
						}
					} else {
						paramTypeName, ok := paramsNames[vl.Name]
						if ok && vl.Text != "" {
							params = append(params, ThreadGroupParams{Type: paramTypeName, Name: "", Value: vl.Text})
						} else if vl.Text != "" {
							params = append(params, ThreadGroupParams{Type: vl.Name, Name: "", Value: vl.Text})
						}
					}
				}
			}
		}
		for i := 0; i < len(tgIndexs); i++ {
			if threadGroupName == tgIndexs[i].Name {
				text := throughputIndexs[i].Name
				for _, v := range thgTimer {
					r := regexp.MustCompile(`\${_{0,2}P?\(?\w*\b,?([\w]+|\d+)\)?}`)
					thgTimerText := r.FindAllStringSubmatch(v.StringProp.Text, -1)
					if len(thgTimerText) > 0 {
						if text == thgTimerText[0][1] {
							value := resParams[text]
							_, err := strconv.Atoi(text)
							if err != nil {
								params = append(params, ThreadGroupParams{Type: "TPS", Name: text, Value: value})
								break
							} else {
								params = append(params, ThreadGroupParams{Type: "TPS", Name: "", Value: text})
								break
							}
						}
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
