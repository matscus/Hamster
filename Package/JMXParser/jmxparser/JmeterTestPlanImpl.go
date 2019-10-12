package jmxparser

import (
	"strings"
)

func (jmx JmeterTestPlan) GetTreadGroupsParams() ([]JMXParserResponse, error) {
	ltg := len(jmx.HashTree.HashTree.ThreadGroup)
	lbztg := len(jmx.HashTree.HashTree.ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup)
	largs := len(jmx.HashTree.HashTree.Arguments.CollectionProp.ElementProp)
	res := make([]JMXParserResponse, 0, ltg+lbztg)
	for i := 0; i < ltg; i++ {
		treadGroupName := jmx.HashTree.HashTree.ThreadGroup[i].Testname
		params := make([]TreadGroupParams, 0, largs)
		for _, vl := range jmx.HashTree.HashTree.ThreadGroup[i].StringProp {
			for _, v := range jmx.HashTree.HashTree.Arguments.CollectionProp.ElementProp {
				if strings.Contains(vl.Text, v.Name) {
					params = append(params, TreadGroupParams{Name: v.Name, Values: v.StringProp[1].Text})
				}
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
			for _, v := range jmx.HashTree.HashTree.Arguments.CollectionProp.ElementProp {
				if strings.Contains(vl.Text, v.Name) {
					params = append(params, TreadGroupParams{Name: v.Name, Values: v.StringProp[1].Text})
				}
			}
		}
		if len(params) > 0 {
			res = append(res, JMXParserResponse{TreadGroupName: treadGroupName, TreadGroupParams: params})
		}
	}
	return res, nil
}
